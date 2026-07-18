-- name: FindTokenByHash :one
SELECT id, token_hash, lobby_id, created_at, expires_at FROM token_auth
WHERE token_hash = sqlc.arg(hash);

-- name: CreateToken :one
INSERT INTO token_auth (token_hash, lobby_id, expires_at)
VALUES (sqlc.arg(token_hash), sqlc.arg(lobby_id), sqlc.arg(expires_at))
RETURNING id, token_hash, lobby_id, created_at, expires_at;

-- name: DeleteTokenByHash :exec
DELETE FROM token_auth WHERE token_hash = sqlc.arg(hash);

-- name: DeleteExpiredTokens :execrows
DELETE FROM token_auth WHERE expires_at < NOW();