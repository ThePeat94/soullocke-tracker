-- name: CreateGameEdition :one
INSERT INTO game_editions (id, name, generation)
VALUES (sqlc.arg(id), sqlc.arg(name), sqlc.arg(generation))
ON CONFLICT DO NOTHING
RETURNING *;

-- name: GetGameEditions :many
SELECT * FROM game_editions;
