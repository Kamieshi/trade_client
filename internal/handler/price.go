package handler

import (
	"context"
	"fmt"

	"github.com/Kamieshi/price_service/protoc"
)

type PriceRPC struct {
	PriceManagerClient protoc.OwnPriceStreamClient
}

func (p *PriceRPC) GetPriceStream(ctx context.Context) (protoc.OwnPriceStream_GetPriceStreamClient, error) {
	stream, err := p.PriceManagerClient.GetPriceStream(ctx, &protoc.GetPriceStreamRequest{})
	if err != nil {
		return nil, fmt.Errorf("handler price / GetPriceStream / Get stream object from RPC server : %v", err)
	}
	return stream, err
}
