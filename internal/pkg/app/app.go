package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rostis232/golibrary/internal/handler"
	"github.com/rostis232/golibrary/internal/postgres"
	"github.com/rostis232/golibrary/internal/service"
)

type App struct {
	Server *echo.Echo
	Handler *handler.Handler
	Service *service.Service
}

func NewApp(pgConfig string) (*App, error) {
	a := App{}
	pg, err := postgres.NewPostgres(pgConfig)
	if err != nil {
		return nil, err
	}
	a.Server = echo.New()
	a.Service = service.NewService(pg)
	a.Handler = handler.NewHandler(a.Service)
	a.Server.Use(middleware.Logger())
	a.Server.Use(middleware.Recover())
	a.Server.Static("/static", "./static")
	a.Server.GET("/", a.Handler.LibraryShow)
	return &a, nil
}

func(a *App) Run(port string) error {
	return a.Server.Start(":"+port)
}