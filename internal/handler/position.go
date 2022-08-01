package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"tradeClient/internal/model"

	"github.com/Kamieshi/position_service/protoc"
)

type PositionRPC struct {
	PositionManagerClient protoc.PositionsManagerClient
}

func (p *PositionRPC) OpenPosition(ctx context.Context, position *model.Position) (string, error) {
	resp, err := p.PositionManagerClient.OpenPosition(ctx, &protoc.OpenPositionRequest{
		Price: &protoc.Price{
			Company: &protoc.Company{
				ID:   position.Company.CompanyID,
				Name: position.Company.Name,
			},
			Ask:  position.Company.Ask,
			Bid:  position.Company.Bid,
			Time: position.Company.Time.Format("2006-01-02T15:04:05.000TZ-07:00"),
		},
		UserID:           position.Client.ID,
		IsSales:          position.IsSales,
		IsFixed:          position.IsFixes,
		MaxProfit:        position.MaxCurrentCost,
		MinProfit:        position.MinCurrentCost,
		CountBuyPosition: position.CountBuyPosition,
	})
	if err != nil {
		return "", fmt.Errorf("handler position / OpenPosition / RPC protocol error : %v", err)
	}
	if resp.Error != "" {
		return "", fmt.Errorf("handler position / OpenPosition / RPC client error : %v", err)
	}
	return resp.ID, nil
}

func (p *PositionRPC) ClosePosition(ctx context.Context, position *model.Position) (int64, error) {
	resp, err := p.PositionManagerClient.ClosePosition(ctx, &protoc.ClosePositionRequest{
		PositionID: position.ID.String(),
		Price: &protoc.Price{
			Company: &protoc.Company{
				ID:   position.Company.CompanyID,
				Name: position.Company.Name,
			},
			Ask:  position.Company.Ask,
			Bid:  position.Company.Bid,
			Time: position.Company.Time.Format("2006-01-02T15:04:05.000TZ-07:00"),
		},
		UserID: position.Client.ID,
	})
	if err != nil {
		return 0, fmt.Errorf("handler position / ClosePosition / RPC protocol error : %v", err)
	}
	if resp.Error != "" {
		return 0, fmt.Errorf("handler position / ClosePosition / RPC client error : %v", err)
	}
	return resp.Profit, nil
}

func (p *PositionRPC) GetPositionByID(ctx context.Context, positionID uuid.UUID) (*model.Position, error) {
	resp, err := p.PositionManagerClient.GetPositionByID(ctx, &protoc.GetPositionByIDRequest{PositionID: positionID.String()})
	if err != nil {
		return nil, fmt.Errorf("handler position / GetPositionByID / RPC protocol error : %v", err)
	}
	if resp.Error != "" {
		return nil, fmt.Errorf("handler position / GetPositionByID / RPC client error : %v", err)
	}
	return &model.Position{
		ID:               resp.Position.PositionID,
		Client:           &model.User{ID: resp.UserID},
		Company:          &model.Price{
			CompanyID: resp.Position.CompanyID,
			Name:      "",
			Ask:       resp.Position.OpenedAsk,
			Bid:       resp.Position.OpenedBid,
			Time:      resp.Position.TimeOpenedPrice,
		},
		IsOpened:         resp.Position.IsOpened,
		PriceClose:       ,
		CurrentCost:      resp.Position.,
		MaxCurrentCost:   0,
		MinCurrentCost:   0,
		CountBuyPosition: 0,
		IsSales:          false,
		IsFixes:          false,
	}, nil
}
