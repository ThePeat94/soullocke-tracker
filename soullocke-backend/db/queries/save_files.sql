-- name: GetAllSaveFiles :many
SELECT * FROM save_files;

-- name: GetSaveFile :one
SELECT * FROM save_files
WHERE id = sqlc.arg(id);