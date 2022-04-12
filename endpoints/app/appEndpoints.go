package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nashmaniac/golang-application-template/adapters"
	v1 "github.com/nashmaniac/golang-application-template/endpoints/app/v1"
)

type appEndPoints struct {
	Usecases adapters.Usecases
	Server   *gin.Engine
}

func NewEndpoints(
	usecases adapters.Usecases,
) (*appEndPoints, error) {

	apiV1, err := v1.V1Api(usecases)
	if err != nil {
		return nil, err
	}
	r := gin.Default()
	
	v1 := r.Group("/v1")
	v1.GET("/healthz", apiV1.Healthz)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return &appEndPoints{
		Usecases: usecases,
		Server:   r,
	}, nil
}
