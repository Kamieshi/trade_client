package echoHTTP

import (
	"net/http"
	"tradeClient/internal/priceStorage"

	"github.com/labstack/echo/v4"
)

type PriceHandler struct {
	PriceStorage *priceStorage.PriceStorage
}

func (p *PriceHandler) GetAll(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "handler price")
}
