package service

import (
	"context"
	"time"

	"github.com/IskanderA1/handly/iternal/domain"
)

type UsersService struct {
}

func NewUsersService() *UsersService {
	return &UsersService{}
}

func (s *UsersService) SignUp(ctx context.Context, input UserSignUpInput) (domain.User, error) {

	user := domain.User{
		Name:         input.Name,
		Password:     input.Password,
		Email:        input.Email,
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Now(),
		RefreshToken: "RefreshToken",
		AccessToken:  "AccessToken",
	}

	return user, nil
}

func (s *UsersService) SignIn(ctx context.Context, input UserSignInInput) (domain.User, error) {
	user := domain.User{
		Name:         "User Name",
		Password:     input.Password,
		Email:        input.Email,
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Now(),
		RefreshToken: "RefreshToken",
		AccessToken:  "AccessToken",
	}

	return user, nil
}

func (s *UsersService) RefreshTokens(ctx context.Context, refreshToken string) (domain.User, error) {
	user := domain.User{
		Name:         "User Name",
		Password:     "Password",
		Email:        "input",
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Now(),
		RefreshToken: "RefreshToken",
		AccessToken:  "AccessToken",
	}

	return user, nil
}
