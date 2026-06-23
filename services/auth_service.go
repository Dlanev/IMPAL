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
	role string,
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

	if role == "" {
		role = "listener" // default role
	}

	user := models.User{
		Username: username,
		Email: email,
		Password: hash,
		Role: role,
	}

	return s.UserRepo.Create(&user)
}

func (s *AuthService) Login(
	email string,
	password string,
) (string, string, error) {

	user, err :=
		s.UserRepo.FindByEmail(email)

	if err != nil {
		return "", "", errors.New(
			"invalid credentials",
		)
	}

	if !utils.CheckPassword(
		password,
		user.Password,
	) {
		return "", "", errors.New(
			"invalid credentials",
		)
	}

	token, err := config.GenerateToken(
		user.ID,
	)

	return token, user.Role, err
}