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
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
	"github.com/sirupsen/logrus"
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
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	company, err := h.services.AdCompany.GetByID(userID, companyIDstring) //TODO probably delete userID?
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, company)
}
func (h *Handler) deleteCompany(c *gin.Context) {
	companyIDstring := c.Param("id")
	h.services.AdCompany.Delete(companyIDstring)
}

func (h *Handler) updateCompany(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	companyIDstring := c.Param("id")
	if err != nil { //TODO!!
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Request.ParseMultipartForm(100000000) // приблизительно 95 мегабайт

	v := c.Request.MultipartForm.Value
	CurrentAmountS := v["CurrentAmount"][0]
	DailyAmountS := v["DailyAmount"][0]
	DaysS := v["Days"][0]
	ca, err := strconv.Atoi(CurrentAmountS)
	da, err := strconv.Atoi(DailyAmountS)
	days, err := strconv.Atoi(DaysS)
	if err != nil {
		logrus.Error(err)
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
		CurrentAmount:          ca,
		DailyAmount:            da,
		Days:                   days,
	}
	logrus.Print(company)

	companyID, err := h.services.AdCompany.Update(company, companyIDstring)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// if companyID == uuid.Nil {
	// 	sendErrorResponse(c, http.StatusInternalServerError, "Company not created")
	// 	return
	// }

	path := "images/" + userID.String() + "/" + companyID.String()

	err = os.MkdirAll(path+storiesFolder, os.ModePerm)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, "Fail creating folder struct for images")
		return
	}
	err = os.MkdirAll(path+postsFolder, os.ModePerm)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, "Fail creating folder struct for images")
		return
	}

	imagesMap := c.Request.MultipartForm.File
	imagesArr := imagesMap["Image"]
	for _, multipartImage := range imagesArr {
		err = writeMultiPartImage(multipartImage, path+storiesFolder)
		if err != nil {
			sendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	imagesSmallArr := imagesMap["ImageSmall"]
	for _, multipartImage := range imagesSmallArr {
		err = writeMultiPartImage(multipartImage, path+postsFolder)
		if err != nil {
			sendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	sendStatusResponse(c, http.StatusOK, company)
}

const (
	storiesFolder = "/stories/"
	postsFolder   = "/posts/"
)

func (h *Handler) createAdCompany(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Request.ParseMultipartForm(100000000) // приблизительно 95 мегабайт

	v := c.Request.MultipartForm.Value
	CurrentAmountS := v["CurrentAmount"][0]
	DailyAmountS := v["DailyAmount"][0]
	DaysS := v["Days"][0]
	ca, err := strconv.Atoi(CurrentAmountS)
	da, err := strconv.Atoi(DailyAmountS)
	days, err := strconv.Atoi(DaysS)
	if err != nil {
		logrus.Error(err)
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
		CurrentAmount:          ca,
		DailyAmount:            da,
		Days:                   days,
	}
	logrus.Print(company)

	companyID, err := h.services.AdCompany.Create(company)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if companyID == uuid.Nil {
		sendErrorResponse(c, http.StatusInternalServerError, "Company not created")
		return
	}

	path := "images/" + userID.String() + "/" + companyID.String()

	err = os.MkdirAll(path+storiesFolder, os.ModePerm)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, "Fail creating folder struct for images")
		return
	}
	err = os.MkdirAll(path+postsFolder, os.ModePerm)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, "Fail creating folder struct for images")
		return
	}

	imagesMap := c.Request.MultipartForm.File
	imagesArr := imagesMap["Image"]
	for _, multipartImage := range imagesArr {
		err = writeMultiPartImage(multipartImage, path+storiesFolder)
		if err != nil {
			sendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	imagesSmallArr := imagesMap["ImageSmall"]
	for _, multipartImage := range imagesSmallArr {
		err = writeMultiPartImage(multipartImage, path+postsFolder)
		if err != nil {
			sendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	sendStatusResponse(c, http.StatusOK, companyID)
}

func writeMultiPartImage(multipartFile *multipart.FileHeader, path string) error {
	file, err := multipartFile.Open()
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}

	out, err := os.Create(fmt.Sprintf("./%s/%s", path, multipartFile.Filename))
	if err != nil {
		return err
	}

	fw := bufio.NewWriter(out)

	fw.Write(buf.Bytes())

	return nil
}
