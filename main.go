package main

import (
	"Isong/config"
	"Isong/controllers"
	"Isong/models"
	"Isong/repositories"
	"Isong/routes"
	"Isong/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	config.ConnectDB()

	config.DB.AutoMigrate(
	&models.User{},
	&models.Song{},
	&models.History{},
	&models.Favorite{},
	&models.Playlist{},
	&models.PlaylistSong{},
	)

	// AUTH

	userRepo := &repositories.UserRepository{}

	authService := &services.AuthService{
		UserRepo: userRepo,
	}

	authController := &controllers.AuthController{
		AuthService: authService,
	}

	// SONGS

	songRepo := &repositories.SongRepository{}

	songService := &services.SongService{
		SongRepo: songRepo,
	}

	songController := &controllers.SongController{
		SongService: songService,
	}

	// HISTORY

	historyRepo :=
	&repositories.HistoryRepository{}

	historyService :=
	&services.HistoryService{
		HistoryRepo: historyRepo,
	}

	historyController :=
	&controllers.HistoryController{
		HistoryService: historyService,
	}

	// FAVOURITE

	favoriteRepo :=
		&repositories.FavoriteRepository{}

	favoriteService :=
		&services.FavoriteService{
		FavoriteRepo: favoriteRepo,
		}

	favoriteController :=
		&controllers.FavoriteController{
		FavoriteService: favoriteService,
		}

	// PLAYLIST

	playlistRepo :=
		&repositories.PlaylistRepository{}

	playlistService :=
		&services.PlaylistService{
		PlaylistRepo: playlistRepo,
		}

	playlistController :=
		&controllers.PlaylistController{
		PlaylistService: playlistService,
		}

	// ROUTER

	r := gin.Default()

	// CORS

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5500",
			"http://127.0.0.1:5500",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
	}))

	// ROUTES

	routes.AuthRoutes(
		r,
		authController,
	)

	routes.SongRoutes(
		r,
		songController,
	)

	routes.HistoryRoutes(
		r,
		historyController,
	)

	routes.FavoriteRoutes(
		r,
		favoriteController,
	)

	routes.PlaylistRoutes(
		r,
		playlistController,
	)

	r.Run(":8080")
}