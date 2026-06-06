package controllers

import (
	"strconv"

	"Isong/services"
	"Isong/utils"

	"github.com/gin-gonic/gin"
)

type SongController struct {
	SongService *services.SongService
}

func (s *SongController) GetAllSongs(
	c *gin.Context,
) {

	songs, err :=
		s.SongService.GetAllSongs()

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
		"songs fetched",
		songs,
	)
}

func (s *SongController) GetSong(
	c *gin.Context,
) {

	idParam := c.Param("id")

	id64, _ := strconv.ParseUint(
		idParam,
		10,
		64,
	)

	song, err :=
		s.SongService.GetSongByID(
			uint(id64),
		)

	if err != nil {

		utils.Error(
			c,
			404,
			"song not found",
		)

		return
	}

	utils.Success(
		c,
		"song found",
		song,
	)
}

func (s *SongController) SearchSongs(
	c *gin.Context,
) {

	query := c.Query("q")

	songs, err :=
		s.SongService.SearchSongs(
			query,
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
		"search results",
		songs,
	)
}

func (s *SongController) TrendingSongs(
	c *gin.Context,
) {

	songs, err :=
		s.SongService.TrendingSongs()

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
		"trending songs",
		songs,
	)
}

