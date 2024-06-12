package handler

import "github.com/rostis232/golibrary/models"

type Handler struct {
	Service Service
}

type Service interface{
	GetAllLibraryItems() ([]models.LibraryItem, error)
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}