package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *apiV1) Healthz(c *gin.Context) {
	output, _ := api.Usecases.GetHealthz(c, "v1")
	c.JSON(http.StatusOK, output)
}
