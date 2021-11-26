package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	logger "github.com/luckyshmo/fb-marketing-app/targetted-back/log"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
)

const (
	storiesFolder = "/stories/"
	postsFolder   = "/posts/"
)

// @Summary get all companies
// @Tags campaign
// @Description get campaign list
// @ID getCampaignList
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {object} []models.AdCampaign
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/campaign/ [get]
func (h *Handler) getCampaignList(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	campaignList, err := h.services.AdCampaign.GetAll(userID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sort.SliceStable(campaignList, func(i, j int) bool {
		return campaignList[i].CreationDate.After(campaignList[j].CreationDate)
	})

	sendStatusResponse(c, http.StatusOK, campaignList)
}

// @Summary get campaign images names
// @Tags campaign
// @Description get campaign images names
// @ID getCampaignImages
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {object} string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/campaign/:id/images/ [get]
func (h *Handler) getCampaignImages(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	campaignIDstring := c.Param("id")

	path := "./images/" + userID.String() + "/" + campaignIDstring

	files, err := ioutil.ReadDir(path + storiesFolder)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	filesSmall, err := ioutil.ReadDir(path + postsFolder)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var names []string
	for _, f := range files {
		names = append(names, storiesFolder+f.Name())
	}
	for _, f := range filesSmall {
		names = append(names, postsFolder+f.Name())
	}

	sendStatusResponse(c, http.StatusOK, names)
}

// @Summary get campaign by id
// @Tags campaign
// @Description get campaign by id
// @ID getCampaignByID
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {object} models.AdCampaign
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/campaign/:id [get]
func (h *Handler) getCampaignByID(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	campaignIDstring := c.Param("id")

	campaign, err := h.services.AdCampaign.GetByID(userID, campaignIDstring) //TODO probably delete userID?
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.services.User.GetById(userID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	campaign.CurrentAmount = user.Amount //TODO ну как бы бред

	sendStatusResponse(c, http.StatusOK, campaign)
}

func (h *Handler) deleteCampaign(c *gin.Context) {
	// id := c.Param("id") //id in request //TODO!

	// uuid, err := uuid.Parse(id)
	// if err != nil {
	// 	sendErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }
	campaignIDstring := c.Param("id")
	if err := h.services.AdCampaign.Delete(campaignIDstring); err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
}

// @Summary update campaign
// @Tags campaign
// @Description update campaign
// @ID updateCampaign
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {object} models.AdCampaign
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/campaign/:id [put]
func (h *Handler) updateCampaign(c *gin.Context) {

	// id := c.Param("id") //id in request //TODO!

	// uuid, err := uuid.Parse(id)
	// if err != nil {
	// 	sendErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }
	campaignIDstring := c.Param("id")
	if campaignIDstring == "" {
		sendErrorResponse(c, http.StatusInternalServerError, "ID is empty")
		return
	}

	campaign, err := parseCampaignFromContext(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	campaignID, err := h.services.AdCampaign.Update(campaign, campaignIDstring)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = createImagesForCampaign(c, campaign.UserId, campaignID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, campaign)
}

// @Summary create campaign
// @Tags campaign
// @Description create campaign
// @ID createCampaign
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {object} uuid.UUID
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/campaign/ [post]
func (h *Handler) createAdCampaign(c *gin.Context) {

	campaign, err := parseCampaignFromContext(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	campaignID, err := h.services.AdCampaign.Create(campaign)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = createImagesForCampaign(c, campaign.UserId, campaignID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, campaignID)
}

func (h *Handler) startAdCampaign(c *gin.Context) {
	id := c.Param("id") //id in request

	uuid, err := uuid.Parse(id)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.AdCampaign.Start(uuid)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//TODO:
	//money service remove payment
	//campaign service change status

	sendStatusResponse(c, http.StatusOK, "Campaign startted")

}

func (h *Handler) stopAdCampaign(c *gin.Context) {
	id := c.Param("id") //id in request

	uuid, err := uuid.Parse(id)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.AdCampaign.Stop(uuid)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, "Campaign stopped")
}

func writeMultiPartImage(multipartFile *multipart.FileHeader, path string) error {
	file, err := multipartFile.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}

	out, err := os.Create(fmt.Sprintf("./%s/%s", path, multipartFile.Filename))
	if err != nil {
		return err
	}
	defer out.Close()

	fw := bufio.NewWriter(out)

	_, err = fw.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func parseCampaignFromContext(c *gin.Context) (models.AdCampaign, error) {
	userID, err := getUserId(c)
	if err != nil {
		return models.AdCampaign{}, fmt.Errorf("get user id form context: %w", err)
	}

	if err = c.Request.ParseMultipartForm(104857600); err != nil { // 100 MB
		return models.AdCampaign{}, fmt.Errorf("parse multiparform: %w", err)
	}

	v := c.Request.MultipartForm.Value
	DailyAmountS := v["DailyAmount"][0]
	DaysS := v["Days"][0]
	da, err := strconv.Atoi(DailyAmountS)
	if err != nil {
		logger.Error(fmt.Errorf("parse daily amount: %w", err))
	}
	days, err := strconv.Atoi(DaysS)
	if err != nil {
		logger.Error(fmt.Errorf("parse days: %w", err))
	}
	campaign := models.AdCampaign{
		UserId:                 userID,
		FbPageId:               v["FbPageId"][0],
		BusinessAddress:        v["BusinessAddress"][0],
		CampaignField:          v["CampaignField"][0],
		CampaignName:           v["CampaignName"][0],
		CompnayPurpose:         v["CompnayPurpose"][0],
		CreativeStatus:         v["CreativeStatus"][0],
		ImagesDescription:      v["ImagesDescription"],
		ImagesSmallDescription: v["ImagesSmallDescription"],
		PostDescription:        v["PostDescription"][0],
		DailyAmount:            da,
		Days:                   days,
	}

	return campaign, nil
}

func createImagesForCampaign(c *gin.Context, userId uuid.UUID, campaignID uuid.UUID) error {
	path := "images/" + userId.String() + "/" + campaignID.String()

	err := os.MkdirAll(path+storiesFolder, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(path+postsFolder, os.ModePerm)
	if err != nil {
		return err
	}

	imagesMap := c.Request.MultipartForm.File
	imagesArr := imagesMap["Image"]
	for _, multipartImage := range imagesArr {
		err = writeMultiPartImage(multipartImage, path+storiesFolder)
		if err != nil {
			return err
		}
	}
	imagesSmallArr := imagesMap["ImageSmall"]
	for _, multipartImage := range imagesSmallArr {
		err = writeMultiPartImage(multipartImage, path+postsFolder)
		if err != nil {
			return err
		}
	}

	return nil
}
