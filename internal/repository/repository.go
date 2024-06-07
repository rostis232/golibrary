package repository

type Library interface{}

type Repository struct {
	Library
}

func NewRepository(lib Library) *Repository{
	return &Repository{
		Library: lib,
	}
}