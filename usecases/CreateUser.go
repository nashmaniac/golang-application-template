package usecases

import (
	"context"

	"github.com/nashmaniac/golang-application-template/models"
)

func (u *usecases) CreateUser(
	ctx context.Context,
	username string,
	password string,
) (*models.User, error) {
	return nil, nil
}
