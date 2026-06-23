package routes

import (
	"Isong/controllers"
	"Isong/middleware"

	"github.com/gin-gonic/gin"
)

func SongRoutes(
	r *gin.Engine,
	song *controllers.SongController,
) {

	api := r.Group("/api/songs")

	api.Use(
		middleware.AuthMiddleware(),
	)

	api.GET(
		"/my-songs",
		middleware.RoleMiddleware("musician"),
		song.GetMySongs,
	)

	api.POST(
		"",
		middleware.RoleMiddleware("musician"),
		song.UploadSong,
	)

	api.GET(
		"",
		song.GetAllSongs,
	)

	api.GET(
		"/trending",
		song.TrendingSongs,
	)

	api.GET(
		"/search",
		song.SearchSongs,
	)

	api.GET(
		"/:id",
		song.GetSong,
	)
}