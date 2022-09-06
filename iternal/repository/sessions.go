package repository

import (
	"context"
	"database/sql"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
	"github.com/google/uuid"
)

type SessionsRepo struct {
	db *db.Queries
}

func NewSessionsRepo(sqlDb *sql.DB) *SessionsRepo {
	return &SessionsRepo{
		db: db.New(sqlDb),
	}
}

func (repo *SessionsRepo) Create(ctx context.Context, param db.CreateSessionParams) (db.Session, error) {
	return repo.db.CreateSession(ctx, param)
}

func (repo *SessionsRepo) GetById(ctx context.Context, uuid uuid.UUID) (db.Session, error) {
	return repo.db.GetSession(ctx, uuid)
}
