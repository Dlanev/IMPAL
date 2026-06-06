package repositories

import (
	"Isong/config"
	"Isong/models"
)

type FavoriteRepository struct{}

func (r *FavoriteRepository) Exists(
	userID uint,
	songID uint,
) bool {

	var count int64

	config.DB.
		Model(&models.Favorite{}).
		Where(
			"user_id = ? AND song_id = ?",
			userID,
			songID,
		).
		Count(&count)

	return count > 0
}

func (r *FavoriteRepository) Create(
	favorite *models.Favorite,
) error {

	return config.DB.
		Create(favorite).
		Error
}

func (r *FavoriteRepository) Delete(
	userID uint,
	songID uint,
) error {

	return config.DB.
		Where(
			"user_id = ? AND song_id = ?",
			userID,
			songID,
		).
		Delete(&models.Favorite{}).
		Error
}

func (r *FavoriteRepository) GetUserFavorites(
	userID uint,
) ([]models.Favorite, error) {

	var favorites []models.Favorite

	err := config.DB.
		Preload("Song").
		Where(
			"user_id = ?",
			userID,
		).
		Order("created_at DESC").
		Find(&favorites).
		Error

	return favorites, err
}

