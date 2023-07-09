package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sadaghiani/concurrent-file-processing/internal/utils"
)

func FixOptionMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func CustomRecovery(c *gin.Context, err any) {
	utils.CreateAbortResponse(c, http.StatusInternalServerError, err)
}
