package echoHTTP

import (
	"net/http"

	"github.com/Kamieshi/trade_client/internal/priceStorage"
	"github.com/labstack/echo/v4"
)

// PositionHandler HTTP echo handler from price
type PriceHandler struct {
	PriceStorage *priceStorage.PriceStorage
}

// GetAll return all prices
// @Tags         price
// @Success      200  {array} model.Price
// @Failure      400  string true "bad Request"
// @Router       /price [get]
// GetAll get user by user Name
func (p *PriceHandler) GetAll(ctx echo.Context) error {
	prices := p.PriceStorage.GetAllPrices()
	return ctx.JSON(http.StatusOK, prices)
}
