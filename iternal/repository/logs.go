package repository

import (
	"context"
	"database/sql"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
)

type LogsRepo struct {
	db *db.Queries
}

func NewLogsRepo(sqlDb *sql.DB) *LogsRepo {
	return &LogsRepo{
		db: db.New(sqlDb),
	}
}

func (repo *LogsRepo) Create(ctx context.Context, param db.CreateLogParams) (db.Log, error) {
	return repo.db.CreateLog(ctx, param)
}

func (repo *LogsRepo) GetById(ctx context.Context, id int64) (db.Log, error) {
	return repo.db.GetLog(ctx, id)
}

func (repo *LogsRepo) GetList(ctx context.Context, param db.ListLogsParams) ([]db.Log, error) {
	return repo.db.ListLogs(ctx, param)
}

func (repo *LogsRepo) Delete(ctx context.Context, id int64) error {
	return repo.db.DeleteLog(ctx, id)
}
