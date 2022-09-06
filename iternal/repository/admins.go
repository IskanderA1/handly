package repository

import (
	"context"
	"database/sql"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
)

type AdminsRepo struct {
	db *db.Queries
}

func NewAdminssRepo(sqlDb *sql.DB) *AdminsRepo {
	return &AdminsRepo{
		db: db.New(sqlDb),
	}
}

func (repo *AdminsRepo) Create(ctx context.Context, param db.CreateAdminParams) (db.Admin, error) {
	return repo.db.CreateAdmin(ctx, param)
}

func (repo *AdminsRepo) GetByUsername(ctx context.Context, username string) (db.Admin, error) {
	return repo.db.GetAdminByUsername(ctx, username)
}

func (repo *AdminsRepo) GetList(ctx context.Context, param db.ListAdminsParams) ([]db.Admin, error) {
	return repo.db.ListAdmins(ctx, param)
}

func (repo *AdminsRepo) Update(ctx context.Context, param db.UpdateAdminParams) (db.Admin, error) {
	return repo.db.UpdateAdmin(ctx, param)
}

func (repo *AdminsRepo) Delete(ctx context.Context, username string) error {
	return repo.db.DeleteAdmin(ctx, username)
}
