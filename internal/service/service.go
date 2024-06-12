package service

import (
	"github.com/rostis232/golibrary/models"
)

type Service struct {
	Repo Repository
}

type Repository interface{
	GetAllLibraryItems()([]models.LibraryItem, error)
}

func NewService(repo Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func(s *Service) GetAllLibraryItems() ([]models.LibraryItem, error) {
	return s.Repo.GetAllLibraryItems()
}