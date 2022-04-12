package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *apiV1) Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "v1",
		"message": "stable",
	})
}
