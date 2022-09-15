package token

import (
	"time"

	"github.com/google/uuid"
)

type AdminPayloadInput struct {
	Username string
	Duration time.Duration
}

type AdminPayload struct {
	ID        uuid.UUID
	Username  string    `json:"name"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewAdminPayload(inp AdminPayloadInput) (*AdminPayload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &AdminPayload{
		ID:        tokenId,
		Username:  inp.Username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(inp.Duration),
	}
	return payload, nil
}

func (p *AdminPayload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
