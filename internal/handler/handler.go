package handler

type Handler struct {
	Service Service
}

type Service interface{}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}