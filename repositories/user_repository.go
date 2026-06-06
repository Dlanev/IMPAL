package repositories

import (
	"Isong/config"
	"Isong/models"
)

type UserRepository struct{}

func (r *UserRepository) Create(
	user *models.User,
) error {

	return config.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(
	email string,
) (*models.User, error) {

	var user models.User

	err := config.DB.
		Where("email = ?", email).
		First(&user).Error

	return &user, err
}

func (r *UserRepository) FindByID(
	id uint,
) (*models.User, error) {

	var user models.User

	err := config.DB.
		First(&user, id).Error

	return &user, err
}