package controllers

import (
	"Isong/services"
	"Isong/utils"

	"github.com/gin-gonic/gin"
)

type HistoryController struct {
	HistoryService *services.HistoryService
}

type PlayRequest struct {
	SongID uint `json:"song_id"`
}

func (h *HistoryController) SavePlay(
	c *gin.Context,
) {

	var req PlayRequest

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
		c.MustGet(
			"user_id",
		).(uint)

	err :=
		h.HistoryService.SavePlay(
			userID,
			req.SongID,
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
		"play recorded",
		nil,
	)
}

func (h *HistoryController) GetHistory(
	c *gin.Context,
) {

	userID :=
		c.MustGet(
			"user_id",
		).(uint)

	histories, err :=
		h.HistoryService.GetHistory(
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
		"history fetched",
		histories,
	)
}

