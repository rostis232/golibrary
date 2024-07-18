package service

import (
	"github.com/rostis232/golibrary/models"
)

type Service struct {
	Repo Repository
}

type Repository interface {
	GetCachedTypes() map[int64]string
	GetCachedLanguages() map[int64]string
	GetCachedDifficulties() map[int64]string
	GetAllLibraryItems() ([]models.LibraryItem, error)
}

func NewService(repo Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) GetAllLibraryItems() ([]models.LibraryItem, error) {
	return s.Repo.GetAllLibraryItems()
}

func (s *Service) GetCachedTypes() map[int64]string {
	return s.Repo.GetCachedTypes()
}

func (s *Service) GetCachedDifficulties() map[int64]string {
	return s.Repo.GetCachedDifficulties()
}

func (s *Service) GetCachedLanguages() map[int64]string {
	return s.Repo.GetCachedLanguages()
}
