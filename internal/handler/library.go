package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rostis232/golibrary/view"
)

func(h Handler) LibraryShow(c echo.Context) error {
	items, err := h.Service.GetAllLibraryItems()
	if err != nil {
		return err
	}
	return render(c, view.ShowLibrary(items))
}
