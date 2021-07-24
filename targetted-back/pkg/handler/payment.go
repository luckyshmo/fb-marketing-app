package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	logger "github.com/luckyshmo/fb-marketing-app/targetted-back/log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
	for i := 0; i < 300; i++ { //TODO! GOVNOKOD
		userID, err := getUserId(c)
		if err != nil {
			sendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		paymentID := c.Param("id")
		url := fmt.Sprintf("https://api.yookassa.ru/v3/payments/%s", paymentID)
		req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))
		if err != nil {
			logger.Error(fmt.Errorf("create request for yookassa: %w", err))
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
			logger.Error(fmt.Errorf("payment status read body: %w", err))
		}

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
	logger.Warn(fmt.Errorf("300 payment ticks passed"))
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
		logger.Error(fmt.Errorf("create token request: %w", err))
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
		logger.Error(fmt.Errorf("read token response body: %w", err))
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
