package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeneralResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func CreateCompleteResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, createGeneralResponse(code, data))
}

func CreateAbortResponse(c *gin.Context, code int, data interface{}) {
	c.AbortWithStatusJSON(code, createGeneralResponse(code, data))
}

func createGeneralResponse(code int, data interface{}) GeneralResponse {
	return GeneralResponse{
		Code:    code,
		Message: http.StatusText(code),
		Data:    data,
	}
}
