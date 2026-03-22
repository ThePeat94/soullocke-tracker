-- name: CreateLobby :one
-- @type LobbyRow
INSERT INTO lobbies (name, password, game_edition_id)
VALUES (sqlc.arg(name), sqlc.arg(password), sqlc.arg(game_edition_id))
RETURNING id, name, game_edition_id;

-- name: GetLobby :one
-- @type LobbyRow
SELECT id, name, game_edition_id FROM lobbies WHERE id = sqlc.arg(id);
