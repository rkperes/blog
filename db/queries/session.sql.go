// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: session.sql

package queries

import (
	"context"
)

const createSession = `-- name: CreateSession :one
insert into
  sessions (id, user_id)
values
  (?, ?) returning id, user_id, created_at
`

type CreateSessionParams struct {
	ID     string
	UserID string
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession, arg.ID, arg.UserID)
	var i Session
	err := row.Scan(&i.ID, &i.UserID, &i.CreatedAt)
	return i, err
}

const getSessionById = `-- name: GetSessionById :one
select
  id, user_id, created_at
from
  sessions
where
  id = ?
`

func (q *Queries) GetSessionById(ctx context.Context, id string) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionById, id)
	var i Session
	err := row.Scan(&i.ID, &i.UserID, &i.CreatedAt)
	return i, err
}
