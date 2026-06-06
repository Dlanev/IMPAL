package controllers

import (

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

