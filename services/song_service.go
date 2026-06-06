package services

import (
	"Isong/models"
	"Isong/repositories"
)

type SongService struct {
	SongRepo *repositories.SongRepository
}

func (s *SongService) GetAllSongs() (
	[]models.Song,
	error,
) {

	return s.SongRepo.GetAll()
}

func (s *SongService) GetSongByID(
	id uint,
) (*models.Song, error) {

	return s.SongRepo.GetByID(id)
}

func (s *SongService) SearchSongs(
	query string,
) ([]models.Song, error) {

	return s.SongRepo.Search(query)
}

func (s *SongService) TrendingSongs() (
	[]models.Song,
	error,
) {

	return s.SongRepo.GetTrending()
}

