package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rostis232/golangjunior/view/user"
)

type LibraryHandler struct{}

func(h LibraryHandler) LibraryShow(c echo.Context) error{
	return render(c, user.Library())
}