// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package tutorial

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTodo = `-- name: CreateTodo :one

INSERT INTO todos (id, title, description, created_at, updated_at)
VALUES ($1,$2,$3,$4,$5)

RETURNING id, created_at, updated_at, title, description
`

type CreateTodoParams struct {
	ID          pgtype.UUID
	Title       string
	Description string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, createTodo,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Description,
	)
	return i, err
}

const getTodo = `-- name: GetTodo :one

SELECT id, created_at, updated_at, title, description FROM todos
WHERE id = $1
`

func (q *Queries) GetTodo(ctx context.Context, id pgtype.UUID) (Todo, error) {
	row := q.db.QueryRow(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Description,
	)
	return i, err
}
