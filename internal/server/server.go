package server

import internal "github.com/gleo-systems/retromeet/internal/server/internal"

type AppServer interface {
	Start()
}

func NewInstance() AppServer {
	return internal.NewAppServer()
}
