package services

import (
	"Isong/models"
	"Isong/repositories"
)

type PlaylistService struct {
	PlaylistRepo *repositories.PlaylistRepository
}

func (s *PlaylistService) CreatePlaylist(
	userID uint,
	name string,
	description string,
) error {

	playlist := models.Playlist{
		UserID: userID,
		Name: name,
		Description: description,
	}

	return s.PlaylistRepo.
		CreatePlaylist(&playlist)
}

func (s *PlaylistService) GetUserPlaylists(
	userID uint,
) ([]models.Playlist, error) {

	return s.PlaylistRepo.
		GetUserPlaylists(userID)
}

func (s *PlaylistService) AddSong(
	playlistID uint,
	songID uint,
) error {

	item := models.PlaylistSong{
		PlaylistID: playlistID,
		SongID: songID,
	}

	return s.PlaylistRepo.
		AddSong(&item)
}

func (s *PlaylistService) RemoveSong(
	playlistID uint,
	songID uint,
) error {

	return s.PlaylistRepo.
		RemoveSong(
			playlistID,
			songID,
		)
}

