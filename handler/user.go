package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rostis232/golangjunior/view/user"
)

type UserHandler struct{}

func(h UserHandler) HandleUserShow(c echo.Context) error{
	return render(c, user.Library())
}