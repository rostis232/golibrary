package main

import (
	"fmt"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/rostis232/golangjunior/handler"
)

func main(){
	app := echo.New()
	libraryHandler := handler.LibraryHandler{}
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Static("/static", "./static")
	app.GET("/library", libraryHandler.LibraryShow)
	app.Start(":6061")

	fmt.Println("working")
}