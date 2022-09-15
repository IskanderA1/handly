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

func (repo *LogsRepo) GetListProjectLog(ctx context.Context, param db.ListProjectLogParams) ([]db.Log, error) {
	return repo.db.ListProjectLog(ctx, param)
}

func (repo *LogsRepo) GetListUserLog(ctx context.Context, param db.ListUserLogParams) ([]db.Log, error) {
	return repo.db.ListUserLog(ctx, param)
}

func (repo *LogsRepo) Delete(ctx context.Context, projectId int64) error {
	return repo.db.DeleteProjectLogs(ctx, projectId)
}
