CREATE TABLE game_editions
(
    id         TEXT PRIMARY KEY,
    name       TEXT        NOT NULL,
    image_src  TEXT,
    deleted_at TIMESTAMPTZ NULL
);

CREATE TABLE lobbies
(
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    password        TEXT NOT NULL,
    game_edition_id TEXT REFERENCES game_editions (id)
);
