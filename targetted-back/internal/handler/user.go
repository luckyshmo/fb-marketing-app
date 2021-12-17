package handler

import (
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Get User By Id
// @Security ApiKeyAuth
// @Tags user
// @Description get user by id
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/:id [get]
func (h *Handler) getUser(c *gin.Context) {

	id := c.Param("id") //id in request

	var uid uuid.UUID
	if id == "0" {
		_uid, err := getUserId(c)
		if err != nil {
			sendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		uid = _uid
	} else {
		_uid, err := uuid.Parse(id)
		if err != nil {
			sendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		uid = _uid
	}

	user, err := h.services.User.GetById(uid)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, user)
}

func (h *Handler) updateBalance(c *gin.Context) {
	am := c.Param("amount")
	id := c.Param("id")

	amount, err := strconv.Atoi(am)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if amount < 0 {
		sendErrorResponse(c, http.StatusBadRequest, "Negative balance")
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.User.SetBalance(uuid, float64(amount)); err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, amount)
}

func (h *Handler) getUserList(c *gin.Context) {

	users, err := h.services.User.GetAll()

	sort.SliceStable(users, func(i, j int) bool {
		return users[i].TimeCreated.After(users[j].TimeCreated)
	})

	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, users)
}

func (h *Handler) getCampaignListByUserID(c *gin.Context) {
	userID := c.Param("id")
	uID, err := uuid.Parse(userID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	campaignList, err := h.services.AdCampaign.GetAll(uID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, campaignList)
}

func (h *Handler) getUserCampaignImages(c *gin.Context) {
	campaignID := c.Param("campaign-id")
	userID := c.Param("id")

	path := "./images/" + userID + "/" + campaignID

	files, err := os.ReadDir(path + storiesFolder)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	filesSmall, err := os.ReadDir(path + postsFolder)
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
