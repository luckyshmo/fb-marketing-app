package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) pendingFacebookPages(c *gin.Context) {
	pendingIDs, err := h.services.FacebookAPI.PageManager.GetPendingPagesID()
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, pendingIDs)
}

func (h *Handler) claimPage(c *gin.Context) {
	fbPageID := c.Param("id")

	message, err := h.services.FacebookAPI.PageManager.PageClaim(fbPageID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, message)
		return
	}

	sendStatusResponse(c, http.StatusOK, message)

}

func (h *Handler) ownedFacebookPages(c *gin.Context) {
	facebookPages, err := h.services.FacebookAPI.PageManager.GetOwnedPages()
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, facebookPages)
}

func (h *Handler) isPageOwnedByID(c *gin.Context) {
	fbPageID := c.Param("id")

	isOwned, err := h.services.FacebookAPI.PageManager.IsPageOwnedByID(fbPageID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if !isOwned {
		sendErrorResponse(c, http.StatusNotFound, "страница не была передана на управление")
		return
	}
	sendStatusResponse(c, http.StatusOK, true)
}

func (h *Handler) deletePendingFacebookPage(c *gin.Context) {

	fbPageID := c.Param("id")

	if err := h.services.FacebookAPI.PageManager.DeletePageByID(fbPageID); err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, fbPageID)
}
