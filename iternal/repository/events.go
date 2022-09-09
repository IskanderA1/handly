package repository

import (
	"context"
	"database/sql"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
)

type EventsRepo struct {
	db *db.Queries
}

func NewEventsRepo(sqlDb *sql.DB) *EventsRepo {
	return &EventsRepo{
		db: db.New(sqlDb),
	}
}

func (repo *EventsRepo) Create(ctx context.Context, param db.CreateEventParams) (db.Event, error) {
	return repo.db.CreateEvent(ctx, param)
}

func (repo *EventsRepo) GetById(ctx context.Context, id int64) (db.Event, error) {
	return repo.db.GetEvent(ctx, id)
}

func (repo *EventsRepo) GetListEventsByProjectId(ctx context.Context, projectId int64) ([]db.Event, error) {
	return repo.db.ListEventsByProjectId(ctx, projectId)
}

func (repo *EventsRepo) Update(ctx context.Context, param db.UpdateEventParams) (db.Event, error) {
	return repo.db.UpdateEvent(ctx, param)
}

func (repo *EventsRepo) Delete(ctx context.Context, id int64) error {
	return repo.db.DeleteEvent(ctx, id)
}
