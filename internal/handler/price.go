// Package handler layer interfaces
package handler

import (
	"context"
	"fmt"

	"github.com/Kamieshi/price_service/protoc"
)

// PriceRPC price rpc interface
type PriceRPC struct {
	PriceManagerClient protoc.OwnPriceStreamClient
}

// GetPriceStream Get price rpc stream
func (p *PriceRPC) GetPriceStream(ctx context.Context) (protoc.OwnPriceStream_GetPriceStreamClient, error) {
	stream, err := p.PriceManagerClient.GetPriceStream(ctx, &protoc.GetPriceStreamRequest{})
	if err != nil {
		return nil, fmt.Errorf("handler price / GetPriceStream / Get stream object from RPC server : %v", err)
	}
	return stream, err
}
