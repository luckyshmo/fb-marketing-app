package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) pendingFacebookPages(c *gin.Context) {
	pendingIDs, err := h.services.Facebook.GetPendingPagesID()
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, pendingIDs)
}

func (h *Handler) deleteFacebookPage(c *gin.Context) {

	fbID := c.Param("id")

	if err := h.services.Facebook.DeletePageByID(fbID); err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	sendStatusResponse(c, http.StatusOK, fbID)
}
