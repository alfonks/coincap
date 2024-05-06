package controller

import (
	"coincap/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (h *HealthCheck) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, entity.HealthCheckResponse{
		Message: "success",
	})
}
