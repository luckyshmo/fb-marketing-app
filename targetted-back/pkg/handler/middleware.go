package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	userId, err := parseHeader(c.GetHeader(authHeader), h.services.Authorization.ParseToken)
	if err != nil {
		sendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	//add userCtx to metadata
	c.Set(userCtx, userId)
}

func parseHeader(header string, parse func(string) (uuid.UUID, error)) (uuid.UUID, error) {
	if header == "" {
		return uuid.Nil, errors.New("empty auth header")
	}

	return parse(header)
}

//return userID from context
func getUserId(c *gin.Context) (uuid.UUID, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return uuid.Nil, errors.New("user not found")
	}

	uuID, ok := id.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("user id is of invalid type")
	}

	return uuID, nil
}
