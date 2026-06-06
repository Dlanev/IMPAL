package routes

import (
	"Isong/controllers"
	"Isong/middleware"

	"github.com/gin-gonic/gin"
)

func HistoryRoutes(
	r *gin.Engine,
	history *controllers.HistoryController,
) {

	api := r.Group("/api/history")

	api.Use(
		middleware.AuthMiddleware(),
	)

	api.POST(
		"",
		history.SavePlay,
	)

	api.GET(
		"",
		history.GetHistory,
	)
}