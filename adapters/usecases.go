package adapters

import (
	"context"

	"github.com/nashmaniac/golang-application-template/models"
)

type Usecases interface {
	GetHealthz(ctx context.Context, version string) (map[string]string, error)
	CreateUser(ctx context.Context, username string, password string) (*models.User, error)
}
