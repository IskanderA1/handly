package domain

import (
	"time"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
)

type Admin struct {
	Username  string    `json:"username"`
	FullName  string    `json:"fullName"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAdmin(admin db.Admin) Admin {
	return Admin{
		Username:  admin.Username,
		FullName:  admin.FullName,
		CreatedAt: admin.CreatedAt,
	}
}
