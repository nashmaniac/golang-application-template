package adapters

import (
	"context"

	"github.com/nashmaniac/golang-application-template/models"
)

type PeristenceStore interface {
	CloseDB(ctx context.Context)
	FindUserByUsername(ctx context.Context, username string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
}
