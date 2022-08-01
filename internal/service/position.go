// Package service trade client
package service

import (
	"context"
	"fmt"
	"tradeClient/internal/handler"
	"tradeClient/internal/model"
	prStor "tradeClient/internal/priceStorage"
)

// Service for work with position
type PositionService struct {
	PositionHandler *handler.PositionRPC
	PriceStorage    *prStor.PriceStorage
}

func (p PositionService) OpenPosition(ctx context.Context, position *model.Position) error {
	err := p.PositionHandler.OpenPosition(ctx, position)
	if err != nil {
		return fmt.Errorf("service position / OpenPosition / open position from handler : %v", err)
	}
	return nil
}

func (p PositionService) ClosePosition(ctx context.Context, position *model.Position) error {
	err := p.PositionHandler.ClosePosition(ctx, position)
	if err != nil {
		return fmt.Errorf("service position / OpenPosition / open position from handler : %v", err)
	}
	return nil
}
