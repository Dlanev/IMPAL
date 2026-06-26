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

	api.GET(
		"/:id/songs",
		playlist.GetPlaylistSongs,
	)

	api.POST(
		"/:id/songs/:songId",
		playlist.AddSong,
	)

	api.DELETE(
		"/:id/songs/:songId",
		playlist.RemoveSong,
	)

	api.DELETE(
		"/:id",
		playlist.DeletePlaylist,
	)
}