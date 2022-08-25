package repository

import (
	"context"
	"database/sql"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Projects interface {
	Create(ctx context.Context, param db.CreateProjectParams) (db.Project, error)
	GetById(ctx context.Context, id int64) (db.Project, error)
	GetList(ctx context.Context, param db.ListProjectsParams) ([]db.Project, error)
	Update(ctx context.Context, param db.UpdateProjectParams) (db.Project, error)
	Delete(ctx context.Context, id int64) error
}

type Accounts interface {
	Create(ctx context.Context, param db.CreateAccountParams) (db.Account, error)
	GetById(ctx context.Context, id string) (db.Account, error)
	GetList(ctx context.Context, param db.ListAccountsParams) ([]db.Account, error)
	Update(ctx context.Context, param db.UpdateAccountParams) (db.Account, error)
	Delete(ctx context.Context, id string) error
}

type Events interface {
	Create(ctx context.Context, param db.CreateEventParams) (db.Event, error)
	GetById(ctx context.Context, id int64) (db.Event, error)
	GetList(ctx context.Context, param db.ListEventsParams) ([]db.Event, error)
	Update(ctx context.Context, param db.UpdateEventParams) (db.Event, error)
	Delete(ctx context.Context, id int64) error
}

type Logs interface {
	Create(ctx context.Context, param db.CreateLogParams) (db.Log, error)
	GetById(ctx context.Context, id int64) (db.Log, error)
	GetList(ctx context.Context, param db.ListLogsParams) ([]db.Log, error)
	Delete(ctx context.Context, id int64) error
}

type Repositories struct {
	Projects Projects
	Accounts Accounts
	Events   Events
	Logs     Logs
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Projects: NewProjectsRepo(db),
		Accounts: NewAccountsRepo(db),
		Events:   NewEventsRepo(db),
		Logs:     NewLogsRepo(db),
	}
}
