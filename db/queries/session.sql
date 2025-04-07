-- name: GetSessionById :one
select
  *
from
  sessions
where
  id = ?;

-- name: CreateSession :one
insert into
  sessions (id, user_id)
values
  (?, ?) returning *;
