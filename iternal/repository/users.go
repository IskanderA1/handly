package repository

import (
	"context"
	"database/sql"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
)

type UsersRepo struct {
	db *db.Queries
}

type NullUser struct {
	User  db.User
	Valid bool
}

func NewUsersRepo(sqlDb *sql.DB) *UsersRepo {
	return &UsersRepo{
		db: db.New(sqlDb),
	}
}

func (repo *UsersRepo) Create(ctx context.Context, param db.CreateUserParams) (db.User, error) {
	return repo.db.CreateUser(ctx, param)
}

func (repo *UsersRepo) GetUserByProjectAccountId(ctx context.Context, accountId string) (NullUser, error) {
	result, err := repo.db.GetUserByProjectAccountId(ctx, sql.NullString{
		String: accountId,
		Valid:  true,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return NullUser{
				User:  db.User{},
				Valid: false,
			}, nil
		}
	}
	return NullUser{
		User:  result,
		Valid: true,
	}, err
}

func (repo *UsersRepo) GetUserByUUID(ctx context.Context, uuid string) (NullUser, error) {
	result, err := repo.db.GetUserByUUID(ctx, sql.NullString{
		String: uuid,
		Valid:  true,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return NullUser{
				User:  db.User{},
				Valid: false,
			}, nil
		}
	}
	return NullUser{
		User:  result,
		Valid: true,
	}, err
}

func (repo *UsersRepo) GetListUsers(ctx context.Context, param db.ListUsersParams) ([]db.User, error) {
	return repo.db.ListUsers(ctx, param)
}

func (repo *UsersRepo) Update(ctx context.Context, param db.UpdateUserParams) (db.User, error) {
	return repo.db.UpdateUser(ctx, param)
}

func (repo *UsersRepo) Delete(ctx context.Context, id string) error {
	return repo.db.DeleteUser(ctx, id)
}
