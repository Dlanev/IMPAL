package services

import (
	"errors"

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

// GetPlaylist returns a playlist only if it belongs to the given user.
func (s *PlaylistService) GetPlaylist(
	playlistID uint,
	userID uint,
) (*models.Playlist, error) {

	playlist, err :=
		s.PlaylistRepo.
			GetPlaylistByID(playlistID)

	if err != nil {
		return nil, errors.New(
			"playlist not found",
		)
	}

	if playlist.UserID != userID {
		return nil, errors.New(
			"unauthorized",
		)
	}

	return playlist, nil
}

func (s *PlaylistService) GetPlaylistSongs(
	playlistID uint,
	userID uint,
) ([]models.PlaylistSong, error) {

	if _, err :=
		s.GetPlaylist(
			playlistID,
			userID,
		); err != nil {

		return nil, err
	}

	return s.PlaylistRepo.
		GetPlaylistSongs(playlistID)
}

func (s *PlaylistService) AddSong(
	userID uint,
	playlistID uint,
	songID uint,
) error {

	if _, err :=
		s.GetPlaylist(
			playlistID,
			userID,
		); err != nil {

		return err
	}

	item := models.PlaylistSong{
		PlaylistID: playlistID,
		SongID: songID,
	}

	return s.PlaylistRepo.
		AddSong(&item)
}

func (s *PlaylistService) RemoveSong(
	userID uint,
	playlistID uint,
	songID uint,
) error {

	if _, err :=
		s.GetPlaylist(
			playlistID,
			userID,
		); err != nil {

		return err
	}

	return s.PlaylistRepo.
		RemoveSong(
			playlistID,
			songID,
		)
}

func (s *PlaylistService) DeletePlaylist(
	userID uint,
	playlistID uint,
) error {

	if _, err :=
		s.GetPlaylist(
			playlistID,
			userID,
		); err != nil {

		return err
	}

	return s.PlaylistRepo.
		DeletePlaylist(playlistID)
}

