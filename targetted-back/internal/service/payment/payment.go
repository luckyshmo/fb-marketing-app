package payment

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	logger "github.com/luckyshmo/fb-marketing-app/targetted-back/log"
)

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

type Youkassa struct {
	client   http.Client
	apiUrl   string
	username string
	password string
}

func NewYoukassaService(cfg config.Youkassa) *Youkassa {
	return &Youkassa{
		client:   http.Client{},
		apiUrl:   cfg.ApiUrl,
		username: cfg.Username,
		password: cfg.Password,
	}
}

func (y *Youkassa) WaitPaymentStatus(paymentID string) (float64, error) {
	for i := 0; i < 300; i++ { //! ticker
		url := fmt.Sprintf(y.apiUrl+"/%s", paymentID)
		req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))
		if err != nil {
			logger.Error(fmt.Errorf("create request for yookassa: %w", err))
		}
		req.Header.Add("Authorization",
			"Basic "+basicAuth(y.username, y.password),
		)

		resp, err := y.client.Do(req)
		if err != nil {
			logger.Error(fmt.Errorf("do GET payment status request: %w", err))
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Error(fmt.Errorf("payment status read body: %w", err))
		}

		var paymentStat paymentStatus
		err = json.Unmarshal(body, &paymentStat)
		if err != nil {
			return 0.0, fmt.Errorf("unmarshal payment response body: %w", err)
		}

		if paymentStat.Status == "succeeded" {
			if paymentAmount, err := strconv.ParseFloat(paymentStat.Amount.Value, 64); err == nil {
				return paymentAmount, nil
			}
		}

		time.Sleep(3 * time.Second)
	}
	return 0.0, errors.New("payment status timeout")
}

func (y *Youkassa) GetPaymentToken(userID uuid.UUID, paymentAmount string) (paymentResponse, error) {
	var tokenRequestJson = []byte(fmt.Sprintf(`{
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
	}`, paymentAmount, userID.String()))
	req, err := http.NewRequest("POST", y.apiUrl, bytes.NewBuffer(tokenRequestJson))
	if err != nil {
		logger.Error(fmt.Errorf("create token request: %w", err))
	}
	req.Header.Set("Idempotence-Key", uuid.New().String())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization",
		"Basic "+basicAuth(y.username, y.password))

	resp, err := y.client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(fmt.Errorf("read token response body: %w", err))
	}
	fmt.Println("response Body:", string(body))

	var paymentResp paymentResponse
	err = json.Unmarshal(body, &paymentResp)
	if err != nil {
		return paymentResponse{}, fmt.Errorf("unmarshal payment response body: %w", err)
	}

	return paymentResp, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
