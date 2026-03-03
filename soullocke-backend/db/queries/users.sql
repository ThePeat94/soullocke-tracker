-- name: GetUser :one
SELECT * FROM users
WHERE id = sqlc.arg(id);

-- name: GetAllUsers :many
SELECT * FROM users;
