package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecases) LoginUser(
	ctx context.Context,
	username string,
	password string,
) (*string, error) {
	user, err := u.PeristenceStore.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("password not matched")
	}

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(24 * time.Hour)).Unix(),
		Issuer:    "golang.application.template",
		Subject:   user.Username,
	}

	signBytes := []byte(u.Configuration.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("could not generate access token. please try again later")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	ss, err := token.SignedString(signBytes)
	if err != nil {
		return nil, err
	}

	return &ss, nil
}
