package routers

import "github.com/gin-gonic/gin"

func (r *Routes) health(rg *gin.RouterGroup) {
	rg.GET("/health", r.controlers.Information.Health)
}
