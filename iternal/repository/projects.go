package repository

import (
	"context"
	"database/sql"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
)

type ProjectsRepo struct {
	db *db.Queries
}

func NewProjectsRepo(sqlDb *sql.DB) *ProjectsRepo {
	return &ProjectsRepo{
		db: db.New(sqlDb),
	}
}

func (repo *ProjectsRepo) Create(ctx context.Context, param db.CreateProjectParams) (db.Project, error) {
	return repo.db.CreateProject(ctx, param)
}

func (repo *ProjectsRepo) GetById(ctx context.Context, id int64) (db.Project, error) {
	return repo.db.GetProject(ctx, id)
}

func (repo *ProjectsRepo) GetList(ctx context.Context, param db.ListProjectsParams) ([]db.Project, error) {
	return repo.db.ListProjects(ctx, param)
}

func (repo *ProjectsRepo) Update(ctx context.Context, param db.UpdateProjectParams) (db.Project, error) {
	return repo.db.UpdateProject(ctx, param)
}

func (repo *ProjectsRepo) Delete(ctx context.Context, id int64) error {
	return repo.db.DeleteProject(ctx, id)
}
