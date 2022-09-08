// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type EventType string

const (
	EventTypeInfo     EventType = "info"
	EventTypeError    EventType = "error"
	EventTypeSuccess  EventType = "success"
	EventTypeUndefind EventType = "undefind"
)

func (e *EventType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EventType(s)
	case string:
		*e = EventType(s)
	default:
		return fmt.Errorf("unsupported scan type for EventType: %T", src)
	}
	return nil
}

type Account struct {
	AccountID    string    `json:"accountID"`
	Name         string    `json:"name"`
	Uuid         string    `json:"uuid"`
	LastUpdateAt time.Time `json:"lastUpdateAt"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Admin struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	FullName  string    `json:"fullName"`
	CreatedAt time.Time `json:"createdAt"`
}

type Event struct {
	ID        int64     `json:"id"`
	ProjectID int64     `json:"projectID"`
	Name      string    `json:"name"`
	EventType EventType `json:"eventType"`
}

type Log struct {
	ID        int64          `json:"id"`
	ProjectID int64          `json:"projectID"`
	EventID   int64          `json:"eventID"`
	AccountID string         `json:"accountID"`
	Data      sql.NullString `json:"data"`
	CreatedAt time.Time      `json:"createdAt"`
}

type Project struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refreshToken"`
	UserAgent    string    `json:"userAgent"`
	ClientIp     string    `json:"clientIp"`
	IsBlocked    bool      `json:"isBlocked"`
	ExpiresAt    time.Time `json:"expiresAt"`
	CreatedAt    time.Time `json:"createdAt"`
}
