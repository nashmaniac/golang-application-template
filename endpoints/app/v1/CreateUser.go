package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nashmaniac/golang-application-template/models"
)

type createUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (v1 *apiV1) CreateUser(
	c *gin.Context,
) {
	var userInput createUserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.AbortWithError(http.StatusBadRequest, &models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid data",
		})
		return
	}

	c.JSON(http.StatusOK, userInput)
}
