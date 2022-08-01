package echoHTTP

import (
	"net/http"
	"tradeClient/internal/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserService *service.UserService
}

func (c *UserHandler) Get(ctx echo.Context) error {
	return nil
}

func (c *UserHandler) GetAll(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "handler user")
}
