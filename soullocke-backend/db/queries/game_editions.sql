-- name: CreateGameEdition :one
INSERT INTO game_editions (id, name, generation)
VALUES (sqlc.arg(id), sqlc.arg(name), sqlc.arg(generation))
ON CONFLICT (id) DO UPDATE SET updated_at = now()
RETURNING *;

-- name: GetGameEditions :many
SELECT * FROM game_editions;

-- name: DeleteGameEdition :one
UPDATE game_editions SET deleted_at = now() WHERE id = sqlc.arg(id)
RETURNING *;
