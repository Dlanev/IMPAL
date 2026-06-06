package services

import (
	"time"

	"Isong/models"
	"Isong/repositories"
)

type HistoryService struct {
	HistoryRepo *repositories.HistoryRepository
}

func (s *HistoryService) SavePlay(
	userID uint,
	songID uint,
) error {

	history := models.History{
		UserID: userID,
		SongID: songID,
		PlayedAt: time.Now(),
	}

	err :=
		s.HistoryRepo.Create(
			&history,
		)

	if err != nil {
		return err
	}

	return s.HistoryRepo.
		IncrementPlayCount(songID)
}

func (s *HistoryService) GetHistory(
	userID uint,
) ([]models.History, error) {

	return s.HistoryRepo.
		GetUserHistory(userID)
}