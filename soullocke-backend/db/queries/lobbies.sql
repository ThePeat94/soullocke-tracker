-- name: CreateLobby :one
INSERT INTO lobbies (name, password, game_edition_id)
VALUES (sqlc.arg(name), sqlc.arg(password), sqlc.arg(game_edition_id))
RETURNING *;

-- name: GetLobby :one
SELECT * FROM lobbies WHERE id = sqlc.arg(id);

-- name: GetLobbies :many
SELECT * FROM lobbies;

-- name: UpdateLobby :one
UPDATE lobbies
SET name = sqlc.arg(name), password = sqlc.arg(password)
WHERE id = sqlc.arg(id)
RETURNING *;