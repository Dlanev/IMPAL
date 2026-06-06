package repositories

import (
	"Isong/config"
	"Isong/models"
)

type SongRepository struct{}

func (r *SongRepository) GetAll() ([]models.Song, error) {

	var songs []models.Song

	err := config.DB.
		Order("id DESC").
		Find(&songs).Error

	return songs, err
}

func (r *SongRepository) GetByID(
	id uint,
) (*models.Song, error) {

	var song models.Song

	err := config.DB.
		First(&song, id).
		Error

	return &song, err
}

func (r *SongRepository) Search(
	query string,
) ([]models.Song, error) {

	var songs []models.Song

	err := config.DB.
		Where(
			"title LIKE ? OR artist LIKE ?",
			"%"+query+"%",
			"%"+query+"%",
		).
		Find(&songs).
		Error

	return songs, err
}

func (r *SongRepository) GetTrending() (
	[]models.Song,
	error,
) {

	var songs []models.Song

	err := config.DB.
		Order("play_count DESC").
		Limit(10).
		Find(&songs).
		Error

	return songs, err
}

