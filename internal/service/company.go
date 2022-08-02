package service

import (
	"fmt"

	"github.com/Kamieshi/trade_client/internal/model"
	"github.com/Kamieshi/trade_client/internal/priceStorage"
)

// CompanyService work with all companies
type CompanyService struct {
	priceStorage *priceStorage.PriceStorage
}

// GetAll companies
func (c *CompanyService) GetAll() ([]model.Price, error) {
	prices := c.priceStorage.GetAllPrices()
	if len(prices) == 0 {
		return nil, fmt.Errorf("service company / GetAll / companies is empty ")
	}
	return prices, nil
}
