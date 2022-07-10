package service

import (
	"context"

	"github.com/IskanderA1/handly/iternal/domain"
)

type UserSignUpInput struct {
	Name     string
	Email    string
	Password string
}

type UserSignInInput struct {
	Email    string
	Password string
}

type Users interface {
	SignUp(ctx context.Context, input UserSignUpInput) (domain.User, error)
	SignIn(ctx context.Context, input UserSignInInput) (domain.User, error)
	RefreshTokens(ctx context.Context, refreshToken string) (domain.User, error)
}

type Services struct {
	Users Users
}

func NewServices() *Services {
	usersService := NewUsersService()

	return &Services{
		Users: usersService,
	}
}
