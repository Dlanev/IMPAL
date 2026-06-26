package controllers

import (

	"strconv"

	"Isong/services"
	"Isong/utils"

	"github.com/gin-gonic/gin"
)

type PlaylistController struct {
	PlaylistService *services.PlaylistService
}

type CreatePlaylistRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func (p *PlaylistController) CreatePlaylist(
	c *gin.Context,
) {

	var req CreatePlaylistRequest

	if err :=
		c.ShouldBindJSON(&req); err != nil {

		utils.Error(
			c,
			400,
			"invalid request",
		)

		return
	}

	userID :=
		c.MustGet("user_id").(uint)

	err :=
		p.PlaylistService.CreatePlaylist(
			userID,
			req.Name,
			req.Description,
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
		"playlist created",
		nil,
	)
}

func (p *PlaylistController) GetPlaylists(
	c *gin.Context,
) {

	userID :=
		c.MustGet("user_id").(uint)

	playlists, err :=
		p.PlaylistService.GetUserPlaylists(
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
		"playlists fetched",
		playlists,
	)
}


func (p *PlaylistController) GetPlaylistSongs(
	c *gin.Context,
) {

	userID :=
		c.MustGet("user_id").(uint)

	playlistID64, _ :=
		strconv.ParseUint(
			c.Param("id"),
			10,
			64,
		)

	songs, err :=
		p.PlaylistService.GetPlaylistSongs(
			uint(playlistID64),
			userID,
		)

	if err != nil {

		utils.Error(
			c,
			404,
			err.Error(),
		)

		return
	}

	utils.Success(
		c,
		"playlist songs fetched",
		songs,
	)
}

func (p *PlaylistController) AddSong(
	c *gin.Context,
) {

	userID :=
		c.MustGet("user_id").(uint)

	playlistID64, _ :=
		strconv.ParseUint(
			c.Param("id"),
			10,
			64,
		)

	songID64, _ :=
		strconv.ParseUint(
			c.Param("songId"),
			10,
			64,
		)

	err :=
		p.PlaylistService.AddSong(
			userID,
			uint(playlistID64),
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
		"song added to playlist",
		nil,
	)
}

func (p *PlaylistController) RemoveSong(
	c *gin.Context,
) {

	userID :=
		c.MustGet("user_id").(uint)

	playlistID64, _ :=
		strconv.ParseUint(
			c.Param("id"),
			10,
			64,
		)

	songID64, _ :=
		strconv.ParseUint(
			c.Param("songId"),
			10,
			64,
		)

	err :=
		p.PlaylistService.RemoveSong(
			userID,
			uint(playlistID64),
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
		"song removed from playlist",
		nil,
	)
}

func (p *PlaylistController) DeletePlaylist(
	c *gin.Context,
) {

	userID :=
		c.MustGet("user_id").(uint)

	playlistID64, _ :=
		strconv.ParseUint(
			c.Param("id"),
			10,
			64,
		)

	err :=
		p.PlaylistService.DeletePlaylist(
			userID,
			uint(playlistID64),
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
		"playlist deleted",
		nil,
	)
}
