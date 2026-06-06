package repositories

import (
	"Isong/config"
	"Isong/models"
)

type PlaylistRepository struct{}

func (r *PlaylistRepository) CreatePlaylist(
	playlist *models.Playlist,
) error {

	return config.DB.
		Create(playlist).
		Error
}

func (r *PlaylistRepository) GetUserPlaylists(
	userID uint,
) ([]models.Playlist, error) {

	var playlists []models.Playlist

	err := config.DB.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&playlists).
		Error

	return playlists, err
}

func (r *PlaylistRepository) GetPlaylistByID(
	id uint,
) (*models.Playlist, error) {

	var playlist models.Playlist

	err := config.DB.
		First(&playlist, id).
		Error

	return &playlist, err
}

func (r *PlaylistRepository) AddSong(
	playlistSong *models.PlaylistSong,
) error {

	return config.DB.
		Create(playlistSong).
		Error
}

func (r *PlaylistRepository) RemoveSong(
	playlistID uint,
	songID uint,
) error {

	return config.DB.
		Where(
			"playlist_id = ? AND song_id = ?",
			playlistID,
			songID,
		).
		Delete(&models.PlaylistSong{}).
		Error
}

func (r *PlaylistRepository) GetPlaylistSongs(
	playlistID uint,
) ([]models.PlaylistSong, error) {

	var songs []models.PlaylistSong

	err := config.DB.
		Preload("Song").
		Where(
			"playlist_id = ?",
			playlistID,
		).
		Find(&songs).
		Error

	return songs, err
}

func (r *PlaylistRepository) DeletePlaylist(
	playlistID uint,
) error {

	return config.DB.
		Delete(
			&models.Playlist{},
			playlistID,
		).
		Error
}

