package app

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/nashmaniac/golang-application-template/adapters"
	"github.com/nashmaniac/golang-application-template/config"
	v1 "github.com/nashmaniac/golang-application-template/endpoints/app/v1"
	"github.com/nashmaniac/golang-application-template/models"
)

type appEndPoints struct {
	Usecases adapters.Usecases
	Server   *gin.Engine
}

func ValidateToken(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationToken := c.Request.Header.Get("authorization")
		tokens := strings.Split(authorizationToken, " ")
		tokenString := tokens[1]
		log.Println(tokenString)

		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method in auth token")
			}
			return []byte(config.SecretKey), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, models.ErrorResponse{
				Message: "unable to parse token",
				Status:  http.StatusForbidden,
			})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok || !token.Valid || claims.Subject == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, models.ErrorResponse{
				Message: "authentication failed",
				Status:  http.StatusForbidden,
			})
			c.Abort()
			return
		}
		log.Println(claims.Subject)
		c.Set("current-user-id", claims.Subject)
		c.Next()
	}
}

func NewEndpoints(
	usecases adapters.Usecases,
	configuration *config.Config,
) (*appEndPoints, error) {

	apiV1, err := v1.V1Api(usecases)
	if err != nil {
		return nil, err
	}
	r := gin.Default()

	v1 := r.Group("/v1")
	unAuthorizedV1 := v1
	unAuthorizedV1.GET("/healthz", apiV1.Healthz)
	unAuthorizedV1.POST("/create-user", apiV1.CreateUser)
	unAuthorizedV1.POST("/login", apiV1.LoginUser)

	authorizedGroupV1 := v1
	authorizedGroupV1.Use(ValidateToken(configuration))
	authorizedGroupV1.GET("/me", apiV1.Me)

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
