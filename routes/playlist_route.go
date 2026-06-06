package routes

import (
	"Isong/controllers"
	"Isong/middleware"

	"github.com/gin-gonic/gin"
)

func PlaylistRoutes(
	r *gin.Engine,
	playlist *controllers.PlaylistController,
) {

	api := r.Group("/api/playlists")

	api.Use(
		middleware.AuthMiddleware(),
	)

	api.POST(
		"",
		playlist.CreatePlaylist,
	)

	api.GET(
		"",
		playlist.GetPlaylists,
	)
}