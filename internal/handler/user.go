package handler

import (
	"context"
	"fmt"

	"tradeClient/internal/model"

	"github.com/Kamieshi/position_service/protoc"
)

type UserRPC struct {
	UserManagerClient protoc.UsersManagerClient
}

func (u *UserRPC) GetByName(ctx context.Context, name string) (*model.User, error) {
	resp, err := u.UserManagerClient.GetUser(ctx, &protoc.GetUserRequest{Name: name})
	if err != nil {
		return nil, fmt.Errorf("handler user / GetByName / get error from RPC server : %v", err)
	}
	return &model.User{
		ID:      resp.User.ID,
		Name:    resp.User.Name,
		Balance: resp.User.Balance,
	}, nil
}

func (u *UserRPC) GetAll(ctx context.Context) ([]*model.User, error) {
	resp, err := u.UserManagerClient.GetAllUsers(ctx, &protoc.GetAllUserRequest{})
	if err != nil {
		return nil, fmt.Errorf("handler user / GetAll / get error from RPC server : %v", err)
	}
	users := make([]*model.User, 0, len(resp.Users))
	for _, userResp := range resp.Users {
		users = append(users, &model.User{
			ID:      userResp.ID,
			Name:    userResp.Name,
			Balance: userResp.Balance,
		})
	}
	return users, nil
}

func (u *UserRPC) UpdateBalance(ctx context.Context, user *model.User, different int64) error {
	resp, err := u.UserManagerClient.AddBalance(ctx, &protoc.AddBalanceRequest{
		UserID:           user.ID,
		DifferentBalance: different,
	})
	if err != nil {
		return fmt.Errorf("handler user / UpdateBalance / RPC protocol error : %v", err)
	}
	if resp.Error != "" {
		return fmt.Errorf("handler user / UpdateBalance / RPC error : %v", resp.Error)
	}
	return nil
}

func (u *UserRPC) CreateUser(ctx context.Context, user *model.User) error {
	resp, err := u.UserManagerClient.CreateUser(ctx, &protoc.CreateUserRequest{User: &protoc.User{
		Name:    user.Name,
		Balance: user.Balance,
	}})
	if err != nil {
		return fmt.Errorf("handler user / CreateUser / RPC protocol error : %v", err)
	}
	if resp.Error != "" {
		return fmt.Errorf("handler user / CreateUser / RPC error: %v", resp.Error)
	}
	return nil
}
