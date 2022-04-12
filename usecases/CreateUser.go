package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/nashmaniac/golang-application-template/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (u *usecases) CreateUser(
	ctx context.Context,
	username string,
	password string,
) (*models.User, error) {

	// first check if there is any user or not with the username
	user, err := u.PeristenceStore.FindUserByUsername(ctx, username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("error fetching user from database")
	}
	if user != nil {
		return nil, fmt.Errorf("user already exists with username " + username)
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("unable to hash password")
	}
	user = &models.User{
		Username: username,
		Password: string(hashedPass),
	}
	user, err = u.PeristenceStore.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
