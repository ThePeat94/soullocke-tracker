-- name: CreateGameEdition :one
INSERT INTO game_editions (id, name)
VALUES (sqlc.arg(id), sqlc.arg(name))
RETURNING *;

-- name: GetGameEditions :many
SELECT * FROM game_editions;

-- name: UpdateImageSrc :exec
UPDATE game_editions SET image_src = sqlc.arg(src) WHERE id = sqlc.arg(id);

-- name: DeleteGameEdition :one
DELETE FROM game_editions WHERE id = sqlc.arg(id)
RETURNING *;
