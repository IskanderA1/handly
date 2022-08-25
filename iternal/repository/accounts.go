package repository

import (
	"context"
	"database/sql"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
)

type AccountsRepo struct {
	db *db.Queries
}

func NewAccountsRepo(sqlDb *sql.DB) *AccountsRepo {
	return &AccountsRepo{
		db: db.New(sqlDb),
	}
}

func (repo *AccountsRepo) Create(ctx context.Context, param db.CreateAccountParams) (db.Account, error) {
	return repo.db.CreateAccount(ctx, param)
}

func (repo *AccountsRepo) GetById(ctx context.Context, id string) (db.Account, error) {
	return repo.db.GetAccount(ctx, id)
}

func (repo *AccountsRepo) GetList(ctx context.Context, param db.ListAccountsParams) ([]db.Account, error) {
	return repo.db.ListAccounts(ctx, param)
}

func (repo *AccountsRepo) Update(ctx context.Context, param db.UpdateAccountParams) (db.Account, error) {
	return repo.db.UpdateAccount(ctx, param)
}

func (repo *AccountsRepo) Delete(ctx context.Context, id string) error {
	return repo.db.DeleteAccount(ctx, id)
}
