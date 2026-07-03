package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health Endpoint
func Health(c *gin.Context) {

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  "UP",
			"service": "Subscriber Service",
			"version": "v2.0.0",
		},
	)
}

// Liveness Endpoint
func Live(c *gin.Context) {

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "ALIVE",
		},
	)
}

// Readiness Endpoint
func Ready(c *gin.Context) {

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  "READY",
			"mongodb": "UP",
			"kafka":   "UP",
			"service": "Subscriber Service",
			"version": "v2.0.0",
		},
	)
}
