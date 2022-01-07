package server

import "github.com/gin-gonic/gin"

func createRouting(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", healthCheckHandler)
	}
}
