package postgres

import (
	"context"

	"github.com/nashmaniac/golang-application-template/models"
)

func (r *postgresRepository) FindUserByUsername(
	ctx context.Context,
	username string,
) (*models.User, error) {
	var user *models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
