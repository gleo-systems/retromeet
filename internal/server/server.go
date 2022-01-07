package server

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppServer interface {
	Start()
}

type appServerImpl struct {
	port int
}

func (as *appServerImpl) Start() {
	fmt.Printf("Starting server on port: %v\n", as.port)
}

func NewInstance() AppServer {
	port := viper.GetInt("port")

	// TODO: akokot - TBD
	return &appServerImpl{port}
}
