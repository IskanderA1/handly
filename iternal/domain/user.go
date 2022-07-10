package domain

import (
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	LastVisitAt  time.Time `json:"lastVisitAt"`
	RegisteredAt time.Time `json:"registeredAt"`
}
