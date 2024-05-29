package main

import (
	"fmt"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/rostis232/golangjunior/handler"
)

func main(){
	app := echo.New()
	userHandler := handler.UserHandler{}
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Static("/static", "./static")
	app.GET("/user", userHandler.HandleUserShow)
	app.Start(":6061")

	fmt.Println("working")
}