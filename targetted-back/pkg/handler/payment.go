package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type paymentRequest struct {
	PaymentAmount string `json:"amount" binding:"required"`
}

func (h *Handler) getPaymentStatus(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	paymentID := c.Param("id")

	amount, err := h.services.Payment.WaitPaymentStatus(paymentID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.User.AddMoney(userID, amount); err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

}

func (h *Handler) getPaymentToken(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var paymentReq paymentRequest
	if err := c.BindJSON(&paymentReq); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	paymentResp, err := h.services.Payment.GetPaymentToken(userID, paymentReq.PaymentAmount)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, paymentResp)
}
