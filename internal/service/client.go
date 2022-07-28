package service

import (
	"context"

	"github.com/Kamieshi/position_service/protoc"
)

type ClientService struct {
	ClientManager protoc.ClientsManagerClient
}

func (c *ClientService) GetAll(ctx context.Context) {
	res, err := c.ClientManager.GetAllClients(ctx, &protoc.GetAllClientRequest{})
}
