package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"Isong/models"
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

func (s *SongController) UploadSong(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Error(c, 401, "unauthorized")
		return
	}

	title := c.PostForm("title")
	artist := c.PostForm("artist")
	album := c.PostForm("album")

	if title == "" || artist == "" {
		utils.Error(c, 400, "title and artist are required")
		return
	}

	// Handle Audio Upload
	audioFile, err := c.FormFile("audio")
	if err != nil {
		utils.Error(c, 400, "audio file is required")
		return
	}

	// Create directories if they don't exist
	os.MkdirAll("uploads/music", os.ModePerm)
	os.MkdirAll("uploads/covers", os.ModePerm)

	audioExt := filepath.Ext(audioFile.Filename)
	audioFilename := fmt.Sprintf("%d_%d%s", userID, time.Now().Unix(), audioExt)
	audioPath := filepath.Join("uploads", "music", audioFilename)

	if err := c.SaveUploadedFile(audioFile, audioPath); err != nil {
		utils.Error(c, 500, "failed to save audio file")
		return
	}

	// Handle Cover Upload (Optional)
	coverPath := ""
	coverFile, err := c.FormFile("cover")
	if err == nil {
		coverExt := filepath.Ext(coverFile.Filename)
		coverFilename := fmt.Sprintf("%d_%d%s", userID, time.Now().Unix(), coverExt)
		fullCoverPath := filepath.Join("uploads", "covers", coverFilename)
		if err := c.SaveUploadedFile(coverFile, fullCoverPath); err == nil {
			coverPath = "/" + filepath.ToSlash(fullCoverPath)
		}
	}

	song := &models.Song{
		UserID:   userID.(uint),
		Title:    title,
		Artist:   artist,
		Album:    album,
		AudioURL: "/" + filepath.ToSlash(audioPath),
		CoverURL: coverPath,
	}

	if err := s.SongService.CreateSong(song); err != nil {
		utils.Error(c, 500, "failed to save song to database")
		return
	}

	utils.Success(c, "song uploaded successfully", song)
}

func (s *SongController) GetMySongs(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Error(c, 401, "unauthorized")
		return
	}

	songs, err := s.SongService.GetMySongs(userID.(uint))
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}

	utils.Success(c, "songs fetched", songs)
}

