CREATE TABLE game_editions
(
    id         TEXT PRIMARY KEY,
    name       TEXT        NOT NULL,
    generation SMALLINT    NOT NULL
);

CREATE TABLE lobbies
(
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            VARCHAR(255) NOT NULL,
    password        TEXT         NOT NULL,
    game_edition_id TEXT REFERENCES game_editions (id),
    created_at      TIMESTAMPTZ      DEFAULT now(),
    updated_at      TIMESTAMPTZ      DEFAULT now(),
    deleted_at      TIMESTAMPTZ  NULL,

    CONSTRAINT name_length CHECK (length(name) BETWEEN 10 AND 255)
);
