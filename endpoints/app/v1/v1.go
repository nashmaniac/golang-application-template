package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nashmaniac/golang-application-template/adapters"
)

type ApiV1 interface {
	Healthz(c *gin.Context)
	CreateUser(c *gin.Context)
	LoginUser(c *gin.Context)
	Me(c *gin.Context)
}

type apiV1 struct {
	Usecases adapters.Usecases
}

func V1Api(
	usecases adapters.Usecases,
) (ApiV1, error) {
	return &apiV1{
		Usecases: usecases,
	}, nil
}
