package service

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
	"github.com/IskanderA1/handly/iternal/domain"
	"github.com/IskanderA1/handly/iternal/repository"
)

type CreateEventInput struct {
	ProjectID int64
	Name      string
	EventType domain.EventType
}

type UpdateEventInput struct {
	ID        int64
	Name      string
	EventType domain.EventType
}

type EventsService struct {
	repository repository.Events
}

func NewEventsService(repository repository.Events) *EventsService {
	return &EventsService{
		repository: repository,
	}
}

func (s *EventsService) Create(ctx context.Context, inp CreateEventInput) (domain.Event, error) {
	param := db.CreateEventParams{
		ProjectID: inp.ProjectID,
		Name:      inp.Name,
		EventType: db.EventType(inp.EventType),
	}

	event, err := s.repository.Create(ctx, param)

	if err != nil {
		return domain.Event{}, err
	}
	return domain.NewEvent(event), err
}

func (s *EventsService) GetById(ctx context.Context, id int64) (domain.Event, error) {

	event, err := s.repository.GetById(ctx, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Event{}, fmt.Errorf("event not found")
		}
		return domain.Event{}, err
	}
	return domain.NewEvent(event), err
}

func (s *EventsService) GetListByProjectId(ctx context.Context, projectID int64) ([]domain.Event, error) {
	res, err := s.repository.GetListEventsByProjectId(ctx, projectID)

	events := make([]domain.Event, 0)

	for _, event := range res {
		events = append(events, domain.NewEvent(event))
	}

	return events, err
}

func (s *EventsService) Update(ctx context.Context, inp UpdateEventInput) (domain.Event, error) {
	event, err := s.repository.Update(ctx, db.UpdateEventParams{
		ID:        inp.ID,
		Name:      inp.Name,
		EventType: db.EventType(inp.EventType),
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Event{}, fmt.Errorf("event not found")
		}
		return domain.Event{}, err
	}
	return domain.NewEvent(event), err
}

func (s *EventsService) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("event not found")
		}
	}
	return err
}
