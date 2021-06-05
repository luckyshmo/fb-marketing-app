package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetCompanyList(c *gin.Context) {
	uuid, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	companyList, err := h.services.AdCompany.GetCompanyList(uuid)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, companyList)
}

func (h *Handler) newCompany(c *gin.Context) {

	uuid, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.Request.ParseMultipartForm(100000000) // приблизительно 95 мегабайт

	v := c.Request.MultipartForm.Value
	company := models.AdCompany{
		UserId:                 uuid,
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

	h.services.AdCompany.CreateCompany(company)

	// files := c.Request.MultipartForm.File
	// filesArr := files["image"]
	// for _, file := range filesArr {
	// 	logrus.Print(file.Filename)
	// 	file1, err := file.Open()
	// 	if err != nil {
	// 		logrus.Error(err)
	// 	}

	// 	buf := bytes.NewBuffer(nil)
	// 	if _, err := io.Copy(buf, file1); err != nil {
	// 		logrus.Error(err)
	// 	}

	// 	// convert []byte to image for saving to file
	// 	img, _, _ := image.Decode(buf)

	// 	//save the imgByte to file
	// 	out, err := os.Create(fmt.Sprintf("./%s", file.Filename)) //сохраняет в '/' контейнера

	// 	if err != nil {
	// 		logrus.Error(err)
	// 	}

	// 	err = png.Encode(out, img)

	// 	if err != nil {
	// 		logrus.Error(err)
	// 	}
	// }
	sendStatusResponse(c, http.StatusOK, http.StatusOK)
}
