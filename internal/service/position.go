// Package service trade client
package service

import (
	"context"
	"fmt"

	"github.com/Kamieshi/trade_client/internal/handler"
	"github.com/Kamieshi/trade_client/internal/model"
	prStor "github.com/Kamieshi/trade_client/internal/priceStorage"
)

// Service for work with position
type PositionService struct {
	PositionHandler *handler.PositionRPC
	PriceStorage    *prStor.PriceStorage
}

func (p *PositionService) OpenPosition(ctx context.Context, position *model.Position) error {
	actualPrice, err := p.PriceStorage.GetPrice(position.Price.CompanyID)
	if err != nil {
		return fmt.Errorf("service position / OpenPosition / Get actual price : %v", err)
	}
	position.Price = &actualPrice
	positionID, err := p.PositionHandler.OpenPosition(ctx, position)
	if err != nil {
		return fmt.Errorf("service position / OpenPosition / open position from handler : %v", err)
	}
	position.ID = positionID
	return nil
}

func (p *PositionService) ClosePosition(ctx context.Context, position *model.Position) (int64, error) {
	profit, err := p.PositionHandler.ClosePosition(ctx, position)
	if err != nil {
		return 0, fmt.Errorf("service position / OpenPosition / open position from handler : %v", err)
	}
	return profit, nil
}

func (p *PositionService) GetAllUserPosition(ctx context.Context, userID string) ([]*model.Position, error) {
	positions, err := p.PositionHandler.GetAllUserPositions(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("serviice position / GetAllUserPosition / Get positions from handler: %v", err)
	}
	return positions, err
}

func (p *PositionService) GetPositionByID(ctx context.Context, positionID string) (*model.Position, error) {
	position, err := p.PositionHandler.GetPositionByID(ctx, positionID)
	if err != nil {
		return nil, fmt.Errorf("serviice position / GetPositionByID / Get position from handler: %v", err)
	}
	return position, nil
}
