package service

import (
	"context"
	"fmt"

	"github.com/Kamieshi/trade_client/internal/handler"
	"github.com/Kamieshi/trade_client/internal/model"
)

// UserService  work with users
type UserService struct {
	UserHandler *handler.UserRPC
}

// GetByName Get user by name
func (c *UserService) GetByName(ctx context.Context, userName string) (*model.User, error) {
	user, err := c.UserHandler.GetByName(ctx, userName)
	if err != nil {
		return nil, fmt.Errorf("service user / Get / get user from handler : %v", err)
	}
	return user, nil
}

// GetAll Get all users
func (c *UserService) GetAll(ctx context.Context) ([]*model.User, error) {
	user, err := c.UserHandler.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service user / GetAll / get users from handler : %v", err)
	}
	return user, nil
}

// UpdateBalance Update balance (positive or negative different)
func (c *UserService) UpdateBalance(ctx context.Context, user *model.User, different int64) error {
	err := c.UserHandler.UpdateBalance(ctx, user, different)
	if err != nil {
		return fmt.Errorf("service user / UpdateBalance /update user balance : %v", err)
	}
	return err
}

// CreateUser create new user
func (c *UserService) CreateUser(ctx context.Context, user *model.User) error {
	err := c.UserHandler.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("service user /  Create User / create user from handler : %v", err)
	}
	return nil
}
