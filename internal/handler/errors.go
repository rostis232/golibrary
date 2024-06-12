package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rostis232/golibrary/view/errors"
)

func (h Handler) NotFoundHandler(c echo.Context) error {
	return render(c, errors.ShowNotFound())
}
