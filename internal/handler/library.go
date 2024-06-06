package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rostis232/golibrary/view/library"
)

func(h Handler) LibraryShow(c echo.Context) error {
	return render(c, library.ShowLibrary())
}
