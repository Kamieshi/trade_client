package handler

import (
	"context"
	"tradeClient/internal/model"

	"github.com/Kamieshi/position_service/protoc"
)

type PositionRPC struct {
	PositionManagerClient protoc.PositionsManagerClient
}

func (p *PositionRPC) OpenPosition(ctx context.Context, position *model.Position) error {
	return nil
}

func (p *PositionRPC) ClosePosition(ctx context.Context, position *model.Position) error {
	return nil
}
