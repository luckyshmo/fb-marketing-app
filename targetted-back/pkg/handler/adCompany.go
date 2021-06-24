package handler

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
	"github.com/sirupsen/logrus"
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

	company.CurrentAmount = int(user.Amount) //TODO ну как бы бред

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
	h.services.AdCompany.Delete(companyIDstring)
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

type paymentRequest struct {
	PaymentAmount string `json:"amount" binding:"required"`
}

type confirmation struct {
	ConfirmationToken string `json:"confirmation_token" binding:"required"`
}

type paymentResponse struct {
	Confirmation confirmation `json:"confirmation" binding:"required"`
	ID           string       `json:"id" binding:"required"`
}

type amount struct {
	Value string `json:"value" binding:"required"`
}

type paymentStatus struct {
	Amount amount `json:"amount" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func (h *Handler) getPaymentStatus(c *gin.Context) {
	for {
		userID, err := getUserId(c)
		if err != nil {
			sendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		paymentID := c.Param("id")
		url := fmt.Sprintf("https://api.yookassa.ru/v3/payments/%s", paymentID)
		req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))
		if err != nil {
			logrus.Error(err)
		}
		req.Header.Add("Authorization", "Basic "+basicAuth("780282", "live_S2yXJKLRFt0Rm_8DMUmwhWb_K46v6YW8QjOsa_FiNew"))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error(err)
		}
		fmt.Println("response Body:", string(body))

		var paymentStat paymentStatus
		err = json.Unmarshal(body, &paymentStat)
		if err != nil {
			sendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if paymentStat.Status == "succeeded" {
			if s, err := strconv.ParseFloat(paymentStat.Amount.Value, 64); err == nil {
				fmt.Println(s)
				h.services.User.AddMoney(userID, s)
				sendStatusResponse(c, http.StatusOK, paymentStat)
				return
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func (h *Handler) getPaymentToken(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	var paymentReq paymentRequest
	if err := c.BindJSON(&paymentReq); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	url := "https://api.yookassa.ru/v3/payments"
	var jsonStr = []byte(fmt.Sprintf(`{
		"amount":
		{
			"value": "%s",
			"currency": "RUB"
		},
		"confirmation":
		{
			"type": "embedded"
		},
		"capture": true,
		"test": true,
		"description": "userID: %s"
	}`, paymentReq.PaymentAmount, userID.String()))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		logrus.Error(err)
	}
	req.Header.Set("Idempotence-Key", uuid.New().String())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth("780282", "live_S2yXJKLRFt0Rm_8DMUmwhWb_K46v6YW8QjOsa_FiNew"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println("response Body:", string(body))

	var paymentResp paymentResponse
	err = json.Unmarshal(body, &paymentResp)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, paymentResp)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
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
		return models.AdCompany{}, err
	}

	c.Request.ParseMultipartForm(104857600) // 100 мегабайт

	v := c.Request.MultipartForm.Value
	DailyAmountS := v["DailyAmount"][0]
	DaysS := v["Days"][0]
	da, err := strconv.Atoi(DailyAmountS)
	if err != nil {
		logrus.Error(err)
	}
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
