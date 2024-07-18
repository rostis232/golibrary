package handler

import "github.com/rostis232/golibrary/models"

type Handler struct {
	Service Service
}

type Service interface {
	GetCachedTypes() map[int64]string
	GetCachedLanguages() map[int64]string
	GetCachedDifficulties() map[int64]string
	GetAllLibraryItems() ([]models.LibraryItem, error)
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}
