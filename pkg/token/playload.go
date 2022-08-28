package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrorExpiredToken = errors.New("Token has expired")
	ErrInvalidToken   = errors.New("Invalid token")
)

type Payload struct {
	ID        uuid.UUID
	Name      string    `json:"name"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(name string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenId,
		Name:      name,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrorExpiredToken
	}
	return nil
}
