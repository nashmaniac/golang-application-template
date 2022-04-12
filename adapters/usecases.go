package adapters

import (
	"context"

	"github.com/nashmaniac/golang-application-template/models"
)

//go:generate mockgen -destination=../mocks/usecases.mock.go -package=mocks --build_flags=--mod=mod github.com/nashmaniac/golang-application-template/adapters Usecases
type Usecases interface {
	GetHealthz(ctx context.Context, version string) (map[string]string, error)
	CreateUser(ctx context.Context, username string, password string) (*models.User, error)
}
