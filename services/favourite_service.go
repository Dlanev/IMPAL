package services

import (
	"errors"

	"Isong/models"
	"Isong/repositories"
)

type FavoriteService struct {
	FavoriteRepo *repositories.FavoriteRepository
}
func (s *FavoriteService) AddFavorite(
	userID uint,
	songID uint,
) error {

	if s.FavoriteRepo.Exists(
		userID,
		songID,
	) {
		return errors.New(
			"song already liked",
		)
	}

	favorite := models.Favorite{
		UserID: userID,
		SongID: songID,
	}

	return s.FavoriteRepo.
		Create(&favorite)
}

func (s *FavoriteService) RemoveFavorite(
	userID uint,
	songID uint,
) error {

	return s.FavoriteRepo.Delete(
		userID,
		songID,
	)
}

func (s *FavoriteService) GetFavorites(
	userID uint,
) ([]models.Favorite, error) {

	return s.FavoriteRepo.
		GetUserFavorites(userID)
}