package http

import (
	"github.com/gin-gonic/gin"
	"github.com/shao1f/user.tag.logic/service"
	"log"
)

var (
	svc *service.Service

	httpServer *gin.Engine
)

func Init(s *service.Service) {
	svc = s

	httpServer = gin.Default()

	initRoute(httpServer)

	go func() {
		if err := httpServer.Run("localhost:9090"); err != nil {
			log.Fatal("http server start failed,err %v", err)
		}
	}()
}

func ShutDown() {
	if svc != nil {
		svc.Close()
	}
}
