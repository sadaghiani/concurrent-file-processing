package server

import (
	"log"

	"github.com/sadaghiani/concurrent-file-processing/internal/routers"
	"github.com/sadaghiani/concurrent-file-processing/pkg/config"
)

func Run(r *routers.Routes) {

	host := config.Config.GetString("host")
	port := config.Config.GetString("port")

	log.Fatalln(r.Run(host + ":" + port))
}
