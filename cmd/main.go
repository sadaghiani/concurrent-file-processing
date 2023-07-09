package main

import (
	"time"

	"github.com/sadaghiani/concurrent-file-processing/internal/controlers"
	"github.com/sadaghiani/concurrent-file-processing/internal/models"
	"github.com/sadaghiani/concurrent-file-processing/internal/repository"
	"github.com/sadaghiani/concurrent-file-processing/internal/routers"
	"github.com/sadaghiani/concurrent-file-processing/internal/server"
	"github.com/sadaghiani/concurrent-file-processing/internal/utils"
	"github.com/sadaghiani/concurrent-file-processing/pkg/config"
	"github.com/sadaghiani/concurrent-file-processing/pkg/database"
	"github.com/sadaghiani/concurrent-file-processing/pkg/logger"
)

//	@title			Concurrent file processing
//	@description	This is a demo version of concurrent file processing
//	@contact.name	sadaghiani
//	@contact.url	https://github.com/sadaghiani
//	@contact.email	sadaghiani.dev@gmail.com
//	@BasePath	    /api/v1

func main() {

	config.NewConfig(utils.CreateMustBindEnvs(), utils.CreateMustBindFlags())
	logger.NewLogger(config.Config.GetInt(utils.MustBindEnvToString(utils.APP_LOG_LEVEL)), time.RFC3339)
	mongoDataStore := database.NewMongoDataStore(logger.Log, utils.CreateMongoDataStoreConfig())
	repository := repository.NewRepository(mongoDataStore)
	models := models.NewModels(repository)
	controlers := controlers.NewControlers(models)
	routers := routers.NewRouter(controlers)
	server.Run(routers)
}
