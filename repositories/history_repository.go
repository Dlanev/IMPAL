package repositories

import (
	"Isong/config"
	"Isong/models"

	"gorm.io/gorm"
)

type HistoryRepository struct{}

func (r *HistoryRepository) Create(
	history *models.History,
) error {

	return config.DB.
		Create(history).
		Error
}

func (r *HistoryRepository) GetUserHistory(
	userID uint,
) ([]models.History, error) {

	var histories []models.History

	err := config.DB.
		Preload("Song").
		Where(
			"user_id = ?",
			userID,
		).
		Order("played_at DESC").
		Limit(50).
		Find(&histories).
		Error

	return histories, err
}

func (r *HistoryRepository) IncrementPlayCount(
	songID uint,
) error {

	return config.DB.
		Model(&models.Song{}).
		Where("id = ?", songID).
		UpdateColumn(
			"play_count",
			gorm.Expr(
				"play_count + 1",
			),
		).
		Error
}

