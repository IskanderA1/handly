// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: projects.sql

package db

import (
	"context"
)

const createProject = `-- name: CreateProject :one
INSERT INTO projects(
    name,
    token
) VALUES (
    $1, $2
) 
RETURNING id, name, token, created_at
`

type CreateProjectParams struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, createProject, arg.Name, arg.Token)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Token,
		&i.CreatedAt,
	)
	return i, err
}

const deleteProject = `-- name: DeleteProject :exec
DELETE FROM projects WHERE id = $1
`

func (q *Queries) DeleteProject(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProject, id)
	return err
}

const getProject = `-- name: GetProject :one
SELECT id, name, token, created_at FROM projects
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProject(ctx context.Context, id int64) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProject, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Token,
		&i.CreatedAt,
	)
	return i, err
}

const listProjects = `-- name: ListProjects :many
SELECT id, name, token, created_at FROM projects
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListProjectsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProjects(ctx context.Context, arg ListProjectsParams) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, listProjects, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Project{}
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Token,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProject = `-- name: UpdateProject :one
UPDATE projects
SET name = $2,
    token = $3
WHERE id = $1
RETURNING id, name, token, created_at
`

type UpdateProjectParams struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, updateProject, arg.ID, arg.Name, arg.Token)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Token,
		&i.CreatedAt,
	)
	return i, err
}