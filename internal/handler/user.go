package handler

import (
	"context"
	"fmt"

	"github.com/Kamieshi/position_service/protoc"
	"github.com/Kamieshi/trade_client/internal/model"
)

// UserRPC user rpc interface
type UserRPC struct {
	UserManagerClient protoc.UsersManagerClient
}

// GetByName get by name
func (u *UserRPC) GetByName(ctx context.Context, name string) (*model.User, error) {
	resp, err := u.UserManagerClient.GetUser(ctx, &protoc.GetUserRequest{Name: name})
	if err != nil {
		return nil, fmt.Errorf("handler user / GetByName / get level gRPC protocol: %v", err)
	}
	if resp.Error != "" {
		return nil, fmt.Errorf("handler user / GetByName / get error from RPC server : %v", resp.Error)
	}
	return &model.User{
		ID:      resp.User.ID,
		Name:    resp.User.Name,
		Balance: resp.User.Balance,
	}, nil
}

// GetAll get all users
func (u *UserRPC) GetAll(ctx context.Context) ([]*model.User, error) {
	resp, err := u.UserManagerClient.GetAllUsers(ctx, &protoc.GetAllUserRequest{})
	if err != nil {
		return nil, fmt.Errorf("handler user / GetAll / get error from RPC server : %v", err)
	}
	if resp.Error != "" {
		return nil, fmt.Errorf("handler user / GetAll / get error from RPC server : %v", resp.Error)
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

// UpdateBalance update balance
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

// CreateUser create new user
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
	user.ID = resp.User.ID
	return nil
}
