package token

import (
	"errors"
)

var (
	ErrExpiredToken = errors.New("Token has expired")
	ErrInvalidToken = errors.New("Invalid token")
)

type Maker[Payload, Input comparable] interface {
	CreateToken(inp Input) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
