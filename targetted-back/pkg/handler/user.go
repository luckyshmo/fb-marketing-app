package handler

import (
	"io/ioutil"
	"net/http"
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

	uuid, err := uuid.Parse(id)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.GetById(uuid)
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

	h.services.User.SetBalance(uuid, float64(amount))

}

func (h *Handler) getUserList(c *gin.Context) {

	users, err := h.services.User.GetAll()

	sort.SliceStable(users, func(i, j int) bool {
		return users[i].DateCreated.After(users[j].DateCreated)
	})

	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, users)
}

func (h *Handler) getCompanyListByUserID(c *gin.Context) {
	userID := c.Param("id")
	uID, err := uuid.Parse(userID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	companyList, err := h.services.AdCompany.GetAll(uID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, companyList)
}

func (h *Handler) getUserCompanyImages(c *gin.Context) {
	companyID := c.Param("company-id")
	userID := c.Param("id")

	path := "./images/" + userID + "/" + companyID

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
