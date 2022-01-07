package server

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gleo-systems/retromeet/internal/log"
	"os"
	"os/signal"
	"time"
)

type AppServerImpl struct {
	config *Config
}

func NewAppServer() *AppServerImpl {
	return &AppServerImpl{readConfiguration()}
}

func (impl *AppServerImpl) Start() {
	address := fmt.Sprint("localhost:", impl.config.Port)
	log.Info("Starting server on address: %s ", address)

	router := gin.New()
	router.Use(loggerAdapter(), recoveryAdapter())

	if err := router.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	createRouting(router)

	go func() {
		if err := router.Run(address); err != nil {
			panic(err)
		}
	}()

	logShutdown()
}

func logShutdown() {
	interrupted := make(chan os.Signal)
	signal.Notify(interrupted, os.Interrupt)
	<-interrupted
	log.Info("Server stopped successfully")
}

func loggerAdapter() gin.HandlerFunc {
	return ginzap.Ginzap(log.GetLogger(), time.RFC3339, true)
}

func recoveryAdapter() gin.HandlerFunc {
	return ginzap.RecoveryWithZap(log.GetLogger(), true)
}
