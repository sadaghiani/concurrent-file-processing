package routers

import (
	"time"

	"github.com/sadaghiani/concurrent-file-processing/internal/controlers"
	"github.com/sadaghiani/concurrent-file-processing/pkg/logger"
	"github.com/sadaghiani/concurrent-file-processing/pkg/middleware"

	"github.com/gin-gonic/gin"

	ginzap "github.com/gin-contrib/zap"
)

type Routes struct {
	Routers    *gin.Engine
	controlers *controlers.Controlers
}

func NewRouter(controlers *controlers.Controlers) *Routes {

	r := Routes{
		Routers:    gin.New(),
		controlers: controlers,
	}

	r.Routers.Use(
		middleware.CORS(),
		middleware.FixOptionMethod(),
		ginzap.Ginzap(logger.Log, time.RFC3339, true),
		ginzap.CustomRecoveryWithZap(logger.Log, true, middleware.CustomRecovery),
	)

	v1 := r.Routers.Group("/api/v1")

	r.root()
	r.docs()
	r.file(v1)
	r.data(v1)
	r.health(v1)

	return &r
}

func (r Routes) Run(addr ...string) error {
	return r.Routers.Run(addr...)
}
