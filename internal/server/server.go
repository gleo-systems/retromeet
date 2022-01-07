package server

import (
	"github.com/gleo-systems/retromeet/internal/env"
	"github.com/gleo-systems/retromeet/internal/log"
)

type AppServer interface {
	Start()
}

type appServerImpl struct {
	port int
}

func (as *appServerImpl) Start() {
	log.Info("Starting server on port: %v", as.port)
}

func NewInstance() AppServer {
	port := env.GetIntVar("port")

	// TODO: akokot - TBD
	return &appServerImpl{port}
}
