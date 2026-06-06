package services

import (
	"errors"

	"Isong/config"
	"Isong/models"
	"Isong/repositories"
	"Isong/utils"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func (s *AuthService) Register(
	username string,
	email string,
	password string,
) error {

	_, err := s.UserRepo.FindByEmail(email)

	if err == nil {
		return errors.New(
			"email already exists",
		)
	}

	hash, err :=
		utils.HashPassword(password)

	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Email: email,
		Password: hash,
	}

	return s.UserRepo.Create(&user)
}

func (s *AuthService) Login(
	email string,
	password string,
) (string, error) {

	user, err :=
		s.UserRepo.FindByEmail(email)

	if err != nil {
		return "", errors.New(
			"invalid credentials",
		)
	}

	if !utils.CheckPassword(
		password,
		user.Password,
	) {
		return "", errors.New(
			"invalid credentials",
		)
	}

	return config.GenerateToken(
		user.ID,
	)
}