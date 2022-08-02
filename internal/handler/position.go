package handler

import (
	"context"
	"fmt"

	"github.com/Kamieshi/position_service/protoc"
	"github.com/Kamieshi/trade_client/internal/model"
)

type PositionRPC struct {
	PositionManagerClient protoc.PositionsManagerClient
}

func (p *PositionRPC) OpenPosition(ctx context.Context, position *model.Position) (string, error) {
	resp, err := p.PositionManagerClient.OpenPosition(ctx, &protoc.OpenPositionRequest{
		Price: &protoc.Price{
			Company: &protoc.Company{
				ID: position.Price.CompanyID,
			},
			Ask:  position.Price.Ask,
			Bid:  position.Price.Bid,
			Time: position.Price.Time,
		},
		UserID:           position.UserID,
		IsSales:          position.IsSales,
		IsFixed:          position.IsFixes,
		MaxProfit:        position.MaxProfit,
		MinProfit:        position.MinProfit,
		CountBuyPosition: position.CountBuyPosition,
	})
	if err != nil {
		return "", fmt.Errorf("handler position / OpenPosition / RPC protocol error : %v", err)
	}
	if resp.Error != "" {
		return "", fmt.Errorf("handler position / OpenPosition / RPC client error : %v", resp.Error)
	}
	return resp.ID, nil
}

func (p *PositionRPC) ClosePosition(ctx context.Context, position *model.Position) (int64, error) {
	resp, err := p.PositionManagerClient.ClosePosition(ctx, &protoc.ClosePositionRequest{
		PositionID: position.ID,
		Price: &protoc.Price{
			Company: &protoc.Company{
				ID:   position.Price.CompanyID,
				Name: position.Price.Name,
			},
			Ask:  position.Price.Ask,
			Bid:  position.Price.Bid,
			Time: position.Price.Time,
		},
		UserID: position.UserID,
	})
	if err != nil {
		return 0, fmt.Errorf("handler position / ClosePosition / RPC protocol error : %v", err)
	}
	if resp.Error != "" {
		return 0, fmt.Errorf("handler position / ClosePosition / RPC client error : %v", resp.Error)
	}
	return resp.Profit, nil
}

func (p *PositionRPC) GetPositionByID(ctx context.Context, positionID string) (*model.Position, error) {
	resp, err := p.PositionManagerClient.GetPositionByID(ctx, &protoc.GetPositionByIDRequest{PositionID: positionID})
	if err != nil {
		return nil, fmt.Errorf("handler position / GetPositionByID / RPC protocol error : %v", err)
	}
	if resp.Error != "" {
		return nil, fmt.Errorf("handler position / GetPositionByID / RPC client error : %v", resp.Error)
	}
	return &model.Position{
		ID:     resp.Position.PositionID,
		UserID: resp.Position.UserID,
		Price: &model.Price{
			CompanyID: resp.Position.CompanyID,
			Ask:       resp.Position.OpenedAsk,
			Bid:       resp.Position.OpenedBid,
			Time:      resp.Position.TimeOpenedPrice,
		},
		IsOpened:         resp.Position.IsOpened,
		Profit:           resp.Position.CloseProfit,
		MaxProfit:        resp.Position.MaxProfit,
		MinProfit:        resp.Position.MinProfit,
		CountBuyPosition: resp.Position.CountPosition,
		IsSales:          resp.Position.IsSale,
		IsFixes:          resp.Position.IsFixed,
	}, nil
}

func (p *PositionRPC) GetAllUserPositions(ctx context.Context, userID string) ([]*model.Position, error) {
	resp, err := p.PositionManagerClient.GetAllUserPositions(ctx, &protoc.GetAllUserPositionsRequest{UserID: userID})
	if err != nil {
		return nil, fmt.Errorf("handler position / GetAllUserPositions / RPC protocol error : %v", err)
	}
	if resp.Error != "" {
		return nil, fmt.Errorf("handler position / GetAllUserPositions / RPC client error : %v", err)
	}
	positions := make([]*model.Position, 0, len(resp.Positions))
	for _, positionProto := range resp.Positions {
		positions = append(positions, &model.Position{
			ID:     positionProto.PositionID,
			UserID: positionProto.UserID,
			Price: &model.Price{
				CompanyID: positionProto.CompanyID,
				Ask:       positionProto.OpenedAsk,
				Bid:       positionProto.OpenedBid,
				Time:      positionProto.TimeOpenedPrice,
			},
			IsOpened:         positionProto.IsOpened,
			Profit:           positionProto.CloseProfit,
			MaxProfit:        positionProto.MaxProfit,
			MinProfit:        positionProto.MinProfit,
			CountBuyPosition: positionProto.CountPosition,
			IsSales:          positionProto.IsSale,
			IsFixes:          positionProto.IsFixed,
		})
	}
	return positions, nil
}
