package routers

import (
	"github.com/sadaghiani/concurrent-file-processing/internal/controlers"

	"github.com/gin-gonic/gin"
)

type data struct{}

func (r *Routes) data(rg *gin.RouterGroup) {

	data := new(data)
	dataRouters := rg.Group("/data")

	data.get(r.controlers, dataRouters)
}

func (f *data) get(c *controlers.Controlers, rg *gin.RouterGroup) {

	data := rg.Group("/")
	data.GET("/", c.Data.Get)
}
