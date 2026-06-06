package controllers

import (
	"strconv"

	"Isong/services"
	"Isong/utils"

	"github.com/gin-gonic/gin"
)

type FavoriteController struct {
	FavoriteService *services.FavoriteService
}

func (f *FavoriteController) AddFavorite(
	c *gin.Context,
) {

	userID :=
		c.MustGet("user_id").(uint)

	songID64, _ :=
		strconv.ParseUint(
			c.Param("songId"),
			10,
			64,
		)

	err :=
		f.FavoriteService.AddFavorite(
			userID,
			uint(songID64),
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
		"song liked",
		nil,
	)
}

func (f *FavoriteController) RemoveFavorite(
	c *gin.Context,
) {

	userID :=
		c.MustGet("user_id").(uint)

	songID64, _ :=
		strconv.ParseUint(
			c.Param("songId"),
			10,
			64,
		)

	err :=
		f.FavoriteService.RemoveFavorite(
			userID,
			uint(songID64),
		)

	if err != nil {

		utils.Error(
			c,
			500,
			err.Error(),
		)

		return
	}

	utils.Success(
		c,
		"song unliked",
		nil,
	)
}

func (f *FavoriteController) GetFavorites(
	c *gin.Context,
) {

	userID :=
		c.MustGet("user_id").(uint)

	favorites, err :=
		f.FavoriteService.GetFavorites(
			userID,
		)

	if err != nil {

		utils.Error(
			c,
			500,
			err.Error(),
		)

		return
	}

	utils.Success(
		c,
		"favorites fetched",
		favorites,
	)
}

