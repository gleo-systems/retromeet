package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Status string

const (
	RUNNING Status = "RUNNING"
	DOWN    Status = "DOWN"
)

type HealthCheck struct {
	Status Status
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, HealthCheck{RUNNING})
}
