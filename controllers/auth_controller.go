package controllers

import (
	"Isong/services"
	"Isong/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *AuthController) Register(
	c *gin.Context,
) {

	var req RegisterRequest

	if err :=
		c.ShouldBindJSON(&req); err != nil {

		utils.Error(
			c,
			400,
			"invalid request",
		)
		return
	}

	err := a.AuthService.Register(
		req.Username,
		req.Email,
		req.Password,
		req.Role,
	)

	if err != nil {

		utils.Error(
			c,
			400,
			err.Error(),
		)
		return
	}

	utils.Success(
		c,
		"registered",
		nil,
	)
}

func (a *AuthController) Login(
	c *gin.Context,
) {

	var req LoginRequest

	if err :=
		c.ShouldBindJSON(&req); err != nil {

		utils.Error(
			c,
			400,
			"invalid request",
		)
		return
	}

	token, role, err :=
		a.AuthService.Login(
			req.Email,
			req.Password,
		)

	if err != nil {

		utils.Error(
			c,
			401,
			err.Error(),
		)
		return
	}

	utils.Success(
		c,
		"login success",
		gin.H{
			"token": token,
			"role":  role,
		},
	)
}

func (a *AuthController) Profile(
	c *gin.Context,
) {

	userID :=
		c.MustGet("user_id").(uint)

	user, err :=
		a.AuthService.
			UserRepo.
			FindByID(userID)

	if err != nil {

		utils.Error(
			c,
			404,
			"user not found",
		)
		return
	}

	user.Password = ""

	utils.Success(
		c,
		"profile",
		user,
	)
}
