package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nashmaniac/golang-application-template/models"
)

type loginUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginUserResponse struct {
	Token string `json:"token"`
}

func (v1 *apiV1) LoginUser(c *gin.Context) {
	var userInput loginUserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.AbortWithError(http.StatusBadRequest, &models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid data",
		})
		return
	}

	token, err := v1.Usecases.LoginUser(c, userInput.Username, userInput.Password)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &models.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, loginUserResponse{
		Token: *token,
	})
}
