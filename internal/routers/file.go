package routers

import (
	"github.com/sadaghiani/concurrent-file-processing/internal/controlers"

	"github.com/gin-gonic/gin"
)

type file struct{}

func (r *Routes) file(rg *gin.RouterGroup) {

	file := new(file)
	fileRouters := rg.Group("/file")

	file.new(r.controlers, fileRouters)
}

func (f *file) new(c *controlers.Controlers, rg *gin.RouterGroup) {

	file := rg.Group("/upload")
	file.POST("/", c.File.Upload)
}
