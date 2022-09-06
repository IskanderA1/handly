package token

import (
	"time"

	"github.com/google/uuid"
)

type ProjectPayloadInput struct {
	ProjectId int64
	Name      string
}

type ProjectPayload struct {
	ID        uuid.UUID
	ProjectId int64     `json:"project_id"`
	Name      string    `json:"name"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewProjectPayload(input ProjectPayloadInput, duration time.Duration) (*ProjectPayload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &ProjectPayload{
		ID:        tokenId,
		ProjectId: input.ProjectId,
		Name:      input.Name,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (p *ProjectPayload) Valid() error {
	if time.Now().After(p.ExpiredAt) && !p.ExpiredAt.IsZero() {
		return ErrExpiredToken
	}
	return nil
}
