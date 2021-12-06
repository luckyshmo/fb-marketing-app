package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func sendErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Warn(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

func sendStatusResponse(c *gin.Context, statusCode int, i interface{}) {
	c.JSON(statusCode, i)
}
