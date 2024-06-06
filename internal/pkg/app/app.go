package app

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/rostis232/golibrary/internal/handler"
	"github.com/rostis232/golibrary/internal/repository"
	"github.com/rostis232/golibrary/internal/service"
)

type App struct {
	Server *echo.Echo
	Handler *handler.Handler
	Service *service.Service
	Repository *repository.Repository
}

func NewApp() *App {
	a := App{}
	a.Server = echo.New()
	a.Repository = repository.NewRepository()
	a.Service = service.NewService(a.Repository)
	a.Handler = handler.NewHandler(a.Service)
	a.Server.Use(middleware.Logger())
	a.Server.Use(middleware.Recover())
	a.Server.Static("/static", "./static")
	a.Server.GET("/library", a.Handler.LibraryShow)
	return &a
}

func(a *App) Run(){
	a.Server.Start(":6061")
}