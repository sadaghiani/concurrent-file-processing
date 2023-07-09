package routers

import (
	_ "github.com/sadaghiani/concurrent-file-processing/internal/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (r *Routes) docs() {
	r.Routers.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
