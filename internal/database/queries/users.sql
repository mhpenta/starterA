-- name: CreateUser :one
INSERT INTO users (
  username,
  email
) VALUES (
  ?, ?
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = ?;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT ? OFFSET ?;

-- name: UpdateUser :one
UPDATE users
SET
  username = ?,
  email = ?,
  updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;