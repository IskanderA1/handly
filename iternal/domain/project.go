package domain

import (
	"time"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
)

type Project struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProjectWithToken struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewProject(project db.Project) Project {
	return Project{
		ID:        project.ID,
		Name:      project.Name,
		CreatedAt: project.CreatedAt,
	}
}

func NewProjectWithToken(project db.Project) ProjectWithToken {
	return ProjectWithToken{
		ID:        project.ID,
		Name:      project.Name,
		Token:     project.Token,
		CreatedAt: project.CreatedAt,
	}
}
