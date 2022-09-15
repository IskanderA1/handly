// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: events.sql

package db

import (
	"context"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO events(
    project_id,
    name,
    event_type
) VALUES (
    $1, $2, $3
) 
RETURNING id, project_id, name, event_type
`

type CreateEventParams struct {
	ProjectID int64     `json:"projectID"`
	Name      string    `json:"name"`
	EventType EventType `json:"eventType"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, createEvent, arg.ProjectID, arg.Name, arg.EventType)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.Name,
		&i.EventType,
	)
	return i, err
}

const deleteEvent = `-- name: DeleteEvent :exec
DELETE FROM events WHERE id = $1
`

func (q *Queries) DeleteEvent(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEvent, id)
	return err
}

const getEvent = `-- name: GetEvent :one
SELECT id, project_id, name, event_type FROM events
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEvent(ctx context.Context, id int64) (Event, error) {
	row := q.db.QueryRowContext(ctx, getEvent, id)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.Name,
		&i.EventType,
	)
	return i, err
}

const getEventByName = `-- name: GetEventByName :one
SELECT id, project_id, name, event_type FROM events
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetEventByName(ctx context.Context, name string) (Event, error) {
	row := q.db.QueryRowContext(ctx, getEventByName, name)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.Name,
		&i.EventType,
	)
	return i, err
}

const listEventsByProjectId = `-- name: ListEventsByProjectId :many
SELECT id, project_id, name, event_type FROM events
WHERE project_id = $1
ORDER BY id
`

func (q *Queries) ListEventsByProjectId(ctx context.Context, projectID int64) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, listEventsByProjectId, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Event{}
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.Name,
			&i.EventType,
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

const updateEvent = `-- name: UpdateEvent :one
UPDATE events
SET name = $2,
    event_type = $3
WHERE id = $1
RETURNING id, project_id, name, event_type
`

type UpdateEventParams struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	EventType EventType `json:"eventType"`
}

func (q *Queries) UpdateEvent(ctx context.Context, arg UpdateEventParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, updateEvent, arg.ID, arg.Name, arg.EventType)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.Name,
		&i.EventType,
	)
	return i, err
}
