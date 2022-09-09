package domain

import db "github.com/IskanderA1/handly/iternal/db/sqlc"

type Event struct {
	ID        int64     `json:"id"`
	ProjectID int64     `json:"projectID"`
	Name      string    `json:"name"`
	EventType EventType `json:"eventType"`
}

type EventType string

const (
	EventTypeInfo     EventType = "info"
	EventTypeError    EventType = "error"
	EventTypeSuccess  EventType = "success"
	EventTypeUndefind EventType = "undefind"
)

func NewEvent(event db.Event) Event {
	return Event{
		ID:        event.ID,
		ProjectID: event.ProjectID,
		Name:      event.Name,
		EventType: EventType(event.EventType),
	}
}

func IsSupportedEvent(eventType string) bool {
	switch eventType {
	case string(EventTypeInfo), string(EventTypeError), string(EventTypeSuccess), string(EventTypeUndefind):
		return true
	}
	return false
}
