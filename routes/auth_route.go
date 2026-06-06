package routes

import (
	"Isong/controllers"
	"Isong/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(
	r *gin.Engine,
	auth *controllers.AuthController,
) {

	api := r.Group("/api/auth")

	api.POST(
		"/register",
		auth.Register,
	)

	api.POST(
		"/login",
		auth.Login,
	)

	api.GET(
		"/profile",
		middleware.AuthMiddleware(),
		auth.Profile,
	)
}