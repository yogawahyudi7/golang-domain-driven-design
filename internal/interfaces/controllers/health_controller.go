package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController handles health check requests
type HealthController struct{}

// NewHealthController creates a new health controller
func NewHealthController() *HealthController {
	return &HealthController{}
}

// Check handles GET /health
func (c *HealthController) Check(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "Service is running",
		"service": "golang-domain-driven-design",
	})
}
