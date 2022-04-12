package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nashmaniac/golang-application-template/models"
)

func (v1 *apiV1) Me(c *gin.Context) {
	currentUser := c.GetString("current-user-id")
	if currentUser == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, &models.ErrorResponse{
			Status:  http.StatusForbidden,
			Message: "forbidden",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"username": currentUser,
	})
}
