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
	companyList, err := h.services.AdCompany.GetCompanyList(userID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, companyList)
}

// func (h *Handler) getCompanyImageByName(c *gin.Context) {
// 	userID, err := getUserId(c)
// 	if err != nil {
// 		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	companyIDstring := c.Param("id")
// 	name := c.Param("name")

// 	path := "./images/" + userID.String() + "/" + companyIDstring

// 	c.Writer.Header().Set("Content-Type", "image/jpeg")

// 	file, err := ioutil.ReadFile(path + "/" + name)
// 	if err != nil {
// 		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	c.Writer.Write(file)

// 	sendStatusResponse(c, http.StatusOK, file)
// }

func (h *Handler) getCompanyImages(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	companyIDstring := c.Param("id")

	path := "./images/" + userID.String() + "/" + companyIDstring

	files, err := ioutil.ReadDir(path)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var names []string
	for _, f := range files {
		names = append(names, f.Name())
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

	company, err := h.services.AdCompany.GetCompanyByID(userID, companyIDstring) //TODO probably delete userID?
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, company)
}

func (h *Handler) createAdCompany(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Request.ParseMultipartForm(100000000) // приблизительно 95 мегабайт

	v := c.Request.MultipartForm.Value
	company := models.AdCompany{
		UserId:                 userID,
		FbPageId:               v["fbPageId"][0],
		BusinessAddress:        v["businessAdress"][0],
		CompanyField:           v["companyField"][0],
		CompanyName:            v["companyName"][0],
		CompnayPurpose:         v["compnayPurpose"][0],
		CreativeStatus:         v["creativeStatus"][0],
		ImagesDescription:      v["imagesDescription"],
		ImagesSmallDescription: v["imagesSmallDescription"],
		PostDescription:        v["postDescription"][0],
	}
	logrus.Print(company)

	companyID, err := h.services.AdCompany.CreateCompany(company)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if companyID == uuid.Nil {
		sendErrorResponse(c, http.StatusInternalServerError, "Company not created")
		return
	}

	path := "images/" + userID.String() + "/" + companyID.String()

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, "Fail creating folder struct for images")
		return
	}

	imagesMap := c.Request.MultipartForm.File
	imagesArr := imagesMap["image"]
	for _, multipartImage := range imagesArr {
		err = writeMultiPartImage(multipartImage, path)
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
