package routes

import (
	"Isong/controllers"
	"Isong/middleware"

	"github.com/gin-gonic/gin"
)

func FavoriteRoutes(
	r *gin.Engine,
	favorite *controllers.FavoriteController,
) {

	api := r.Group("/api/favorites")

	api.Use(
		middleware.AuthMiddleware(),
	)

	api.GET(
		"",
		favorite.GetFavorites,
	)

	api.POST(
		"/:songId",
		favorite.AddFavorite,
	)

	api.DELETE(
		"/:songId",
		favorite.RemoveFavorite,
	)
}