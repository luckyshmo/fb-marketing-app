package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service"
	logger "github.com/luckyshmo/fb-marketing-app/targetted-back/log"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
)

const (
	storiesFolder = "/stories/"
	postsFolder   = "/posts/"
)

// @Summary get all campaigns
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
		return campaignList[i].TimeCreated.After(campaignList[j].TimeCreated)
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

	files, err := os.ReadDir(path + storiesFolder)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	filesSmall, err := os.ReadDir(path + postsFolder)
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

	// user, err := h.services.User.GetById(userID)
	// if err != nil {
	// 	sendErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

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

	imgArr, err := parseImgFromReq(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//! should return proper campaign
	_, err = h.services.AdCampaign.Update(campaign, campaignIDstring, imgArr)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, "")
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

	imgArr, err := parseImgFromReq(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	campaignID, err := h.services.AdCampaign.Create(campaign, imgArr)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, campaignID)
}

func parseImgFromReq(c *gin.Context) (img []service.ImageWithMessage, err error) {
	imagesMap := c.Request.MultipartForm.File
	imagesArr := imagesMap["Image"]
	for _, multipartImage := range imagesArr {
		file, openErr := multipartImage.Open()
		if err != nil {
			err = openErr
			return
		}
		defer file.Close()

		image, readErr := io.ReadAll(file)
		if err != nil {
			err = readErr
			return
		}
		img = append(img, service.ImageWithMessage{
			Img:     image,
			IsStory: true,
			Message: "", //!
		})
	}
	imagesSmallArr := imagesMap["ImageSmall"]
	for _, multipartImage := range imagesSmallArr {
		file, openErr := multipartImage.Open()
		if err != nil {
			err = openErr
			return
		}
		defer file.Close()

		image, readErr := io.ReadAll(file)
		if err != nil {
			err = readErr
			return
		}
		img = append(img, service.ImageWithMessage{
			Img:     image,
			IsStory: false,
			Message: "", //!
		})
	}
	return
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

func parseCampaignFromContext(c *gin.Context) (models.AdCampaign, error) {
	userID, err := getUserId(c)
	if err != nil {
		return models.AdCampaign{}, fmt.Errorf("get user id form context: %w", err)
	}

	if err = c.Request.ParseMultipartForm(104857600); err != nil { // 100 MB
		return models.AdCampaign{}, fmt.Errorf("parse multiparform: %w", err)
	}

	AdCampaign := c.Request.MultipartForm.Value

	dailyBudgetS := AdCampaign["DailyBudget"][0]
	dailyBudget, err := strconv.ParseFloat(dailyBudgetS, 64)
	if err != nil {
		logger.Error(fmt.Errorf("parse daily campaign budget: %w", err))
	}

	durationS := AdCampaign["Duration"][0]
	duration, err := strconv.Atoi(durationS)
	if err != nil {
		logger.Error(fmt.Errorf("parse duration: %w", err))
	}

	campaign := models.AdCampaign{
		UserId:          userID,
		FbPageId:        AdCampaign["FbPageId"][0],
		BusinessAddress: AdCampaign["BusinessAddress"][0],
		Field:           AdCampaign["Field"][0],
		Name:            AdCampaign["Name"][0],
		Objective:       AdCampaign["Objective"][0],
		CreativeStatus:  AdCampaign["CreativeStatus"][0],
		DailyBudget:     dailyBudget,
		Duration:        duration,
	} //! not all fields!

	return campaign, nil
}

//! Deprecated but probably useful
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

//! Deprecated but probably useful
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
