package token

import (
	"errors"
	"time"
)

var (
	ErrExpiredToken = errors.New("Token has expired")
	ErrInvalidToken = errors.New("Invalid token")
)

type Maker interface {
	CreateProjectToken(input ProjectPayloadInput) (string, *ProjectPayload, error)
	VerifyProjectToken(token string) (*ProjectPayload, error)
	CreateAdminToken(username string, duration time.Duration) (string, *AdminPayload, error)
	VerifyAdminToken(token string) (*AdminPayload, error)
}
