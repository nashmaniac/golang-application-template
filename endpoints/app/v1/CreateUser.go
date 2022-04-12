package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nashmaniac/golang-application-template/models"
)

type createUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type createUserResponse struct {
	Username  string    `json:"username"`
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
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

	user, err := v1.Usecases.CreateUser(c, userInput.Username, userInput.Password)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &models.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, createUserResponse{
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		ID:        user.ID,
	})
}
