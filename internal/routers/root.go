package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Routes) root() {
	r.Routers.GET("/", func(ctx *gin.Context) { ctx.Redirect(http.StatusMovedPermanently, "/docs/index.html") })

}
