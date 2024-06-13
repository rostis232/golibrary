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
	types := h.Service.GetCachedTypes()
	diff := h.Service.GetCachedDifficulties()
	lang := h.Service.GetCachedLanguages()
	return render(c, view.ShowLibrary(items, types, diff, lang))
}
