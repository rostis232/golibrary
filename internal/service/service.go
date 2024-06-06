package service

type Service struct {
	Repo Repository
}

type Repository interface{}

func NewService(repo Repository) *Service {
	return &Service{
		Repo: repo,
	}
}