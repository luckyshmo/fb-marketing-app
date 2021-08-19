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

func (h *Handler) getCompanyList(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	companyList, err := h.services.AdCompany.GetAll(userID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sort.SliceStable(companyList, func(i, j int) bool {
		return companyList[i].CreationDate.After(companyList[j].CreationDate)
	})

	sendStatusResponse(c, http.StatusOK, companyList)
}

func (h *Handler) getCompanyImages(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	companyIDstring := c.Param("id")

	path := "./images/" + userID.String() + "/" + companyIDstring

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

func (h *Handler) getCompanyByID(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	companyIDstring := c.Param("id")

	company, err := h.services.AdCompany.GetByID(userID, companyIDstring) //TODO probably delete userID?
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.services.User.GetById(userID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	company.CurrentAmount = user.Amount //TODO ну как бы бред

	sendStatusResponse(c, http.StatusOK, company)
}

func (h *Handler) deleteCompany(c *gin.Context) {
	// id := c.Param("id") //id in request //TODO!

	// uuid, err := uuid.Parse(id)
	// if err != nil {
	// 	sendErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }
	companyIDstring := c.Param("id")
	if err := h.services.AdCompany.Delete(companyIDstring); err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
}

func (h *Handler) updateCompany(c *gin.Context) {

	// id := c.Param("id") //id in request //TODO!

	// uuid, err := uuid.Parse(id)
	// if err != nil {
	// 	sendErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }
	companyIDstring := c.Param("id")
	if companyIDstring == "" {
		sendErrorResponse(c, http.StatusInternalServerError, "ID is empty")
		return
	}

	company, err := parseCompanyFromContext(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	companyID, err := h.services.AdCompany.Update(company, companyIDstring)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = createImagesForCompany(c, company.UserId, companyID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, company)
}

func (h *Handler) createAdCompany(c *gin.Context) {

	company, err := parseCompanyFromContext(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	companyID, err := h.services.AdCompany.Create(company)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = createImagesForCompany(c, company.UserId, companyID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, companyID)
}

func (h *Handler) startAdCompany(c *gin.Context) {
	id := c.Param("id") //id in request

	uuid, err := uuid.Parse(id)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.AdCompany.Start(uuid)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//TODO:
	//money service remove payment
	//company service change status

	sendStatusResponse(c, http.StatusOK, "Company startted")

}

func (h *Handler) stopAdCompany(c *gin.Context) {
	id := c.Param("id") //id in request

	uuid, err := uuid.Parse(id)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.AdCompany.Stop(uuid)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, "Company stopped")
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

func parseCompanyFromContext(c *gin.Context) (models.AdCompany, error) {
	userID, err := getUserId(c)
	if err != nil {
		return models.AdCompany{}, fmt.Errorf("get user id form context: %w", err)
	}

	if err = c.Request.ParseMultipartForm(104857600); err != nil { // 100 MB
		return models.AdCompany{}, fmt.Errorf("parse multiparform: %w", err)
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
	company := models.AdCompany{
		UserId:                 userID,
		FbPageId:               v["FbPageId"][0],
		BusinessAddress:        v["BusinessAddress"][0],
		CompanyField:           v["CompanyField"][0],
		CompanyName:            v["CompanyName"][0],
		CompnayPurpose:         v["CompnayPurpose"][0],
		CreativeStatus:         v["CreativeStatus"][0],
		ImagesDescription:      v["ImagesDescription"],
		ImagesSmallDescription: v["ImagesSmallDescription"],
		PostDescription:        v["PostDescription"][0],
		DailyAmount:            da,
		Days:                   days,
	}

	return company, nil
}

func createImagesForCompany(c *gin.Context, userId uuid.UUID, companyID uuid.UUID) error {
	path := "images/" + userId.String() + "/" + companyID.String()

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
