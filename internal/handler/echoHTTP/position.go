package echoHTTP

import (
	"net/http"
	"tradeClient/internal/service"

	"github.com/labstack/echo/v4"
)

type PositionHandler struct {
	PositionService *service.PositionService
}

func (*PositionHandler) GetAll(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "handler position")
}
