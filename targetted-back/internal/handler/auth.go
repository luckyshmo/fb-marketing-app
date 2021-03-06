package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
)

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User

		if err := c.BindJSON(&input); err != nil {
			sendErrorResponse(c, http.StatusBadRequest, "invalid input body")
			return
		}

		id, err := h.services.Authorization.CreateUser(input)
		if err != nil {
			sendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		sendStatusResponse(c, http.StatusOK, map[string]interface{}{
			"id": id, //JSON body
		})
	}
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input signInInput

		if err := c.BindJSON(&input); err != nil {
			sendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
		if err != nil {
			sendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		sendStatusResponse(c, http.StatusOK, map[string]interface{}{
			"token": token, //JSON body
			"user": models.User{
				Email: input.Email,
			},
		})
	}
}
