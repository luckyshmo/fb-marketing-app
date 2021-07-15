package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	logger "github.com/luckyshmo/fb-marketing-app/targetted-back/log"
)

type errorResponse struct {
	Message string `json:"message"`
}

func sendErrorResponse(c *gin.Context, statusCode int, message string) {
	if statusCode > 401 { //Proper check what we need to log
		logger.Error(fmt.Errorf(message))
	}
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

func sendStatusResponse(c *gin.Context, statusCode int, i interface{}) {
	c.JSON(statusCode, i)
}
