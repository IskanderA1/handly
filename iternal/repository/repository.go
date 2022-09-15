package repository

import (
	"context"
	"database/sql"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
	"github.com/google/uuid"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Admins interface {
	Create(ctx context.Context, param db.CreateAdminParams) (db.Admin, error)
	GetByUsername(ctx context.Context, username string) (db.Admin, error)
	GetList(ctx context.Context, param db.ListAdminsParams) ([]db.Admin, error)
	Update(ctx context.Context, param db.UpdateAdminParams) (db.Admin, error)
	Delete(ctx context.Context, username string) error
}

type Projects interface {
	Create(ctx context.Context, param db.CreateProjectParams) (db.Project, error)
	GetById(ctx context.Context, id int64) (db.Project, error)
	GetList(ctx context.Context, param db.ListProjectsParams) ([]db.Project, error)
	Update(ctx context.Context, param db.UpdateProjectParams) (db.Project, error)
	Delete(ctx context.Context, id int64) error
}

type Users interface {
	Create(ctx context.Context, param db.CreateUserParams) (db.User, error)
	GetUserByProjectAccountId(ctx context.Context, accountId string) (NullUser, error)
	GetUserByUUID(ctx context.Context, uuid string) (NullUser, error)
	GetListUsers(ctx context.Context, param db.ListUsersParams) ([]db.User, error)
	Update(ctx context.Context, param db.UpdateUserParams) (db.User, error)
	Delete(ctx context.Context, id string) error
}

type Events interface {
	Create(ctx context.Context, param db.CreateEventParams) (db.Event, error)
	GetById(ctx context.Context, id int64) (db.Event, error)
	GetByName(ctx context.Context, name string) (db.Event, error)
	GetListEventsByProjectId(ctx context.Context, projectID int64) ([]db.Event, error)
	Update(ctx context.Context, param db.UpdateEventParams) (db.Event, error)
	Delete(ctx context.Context, id int64) error
}

type Logs interface {
	Create(ctx context.Context, param db.CreateLogParams) (db.Log, error)
	GetListProjectLog(ctx context.Context, param db.ListProjectLogParams) ([]db.Log, error)
	GetListUserLog(ctx context.Context, param db.ListUserLogParams) ([]db.Log, error)
	Delete(ctx context.Context, id int64) error
}

type Sessions interface {
	Create(ctx context.Context, param db.CreateSessionParams) (db.Session, error)
	GetById(ctx context.Context, uuid uuid.UUID) (db.Session, error)
}

type Repositories struct {
	Admins   Admins
	Projects Projects
	Users    Users
	Events   Events
	Logs     Logs
	Sessions Sessions
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Admins:   NewAdminssRepo(db),
		Projects: NewProjectsRepo(db),
		Users:    NewUsersRepo(db),
		Events:   NewEventsRepo(db),
		Logs:     NewLogsRepo(db),
		Sessions: NewSessionsRepo(db),
	}
}
