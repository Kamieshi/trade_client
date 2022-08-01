package handler

import (
	"context"
	"tradeClient/internal/model"

	"github.com/Kamieshi/position_service/protoc"
)

type UserRPC struct {
	UserManagerClient protoc.ClientsManagerClient
}

func (u *UserRPC) GetByName(ctx context.Context, name string) (*model.User, error) {
	return nil, nil
}

func (u *UserRPC) GetAll(ctx context.Context) ([]*model.User, error) {
	return nil, nil
}

func (u *UserRPC) UpdateBalance(ctx context.Context, user *model.User, different int64) error {
	return nil
}

func (u *UserRPC) CreateUser(ctx context.Context, user *model.User) error {
	return nil
}
