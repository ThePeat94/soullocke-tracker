CREATE TABLE game_editions
(
    id         INT PRIMARY KEY,
    code_name  TEXT     NOT NULL,
    generation SMALLINT NOT NULL
);

CREATE TABLE lobbies
(
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            VARCHAR(255)                      NOT NULL,
    password        TEXT                              NOT NULL,
    game_edition_id INT REFERENCES game_editions (id) NOT NULL,
    created_at      TIMESTAMPTZ      DEFAULT now(),
    updated_at      TIMESTAMPTZ      DEFAULT now(),
    deleted_at      TIMESTAMPTZ                       NULL,

    CONSTRAINT name_length CHECK (length(name) BETWEEN 10 AND 255)
);

CREATE TABLE languages
(
    id      INT PRIMARY KEY,
    iso639  VARCHAR(10)  NOT NULL, -- language code
    iso3166 VARCHAR(2)   NOT NULL, -- country code
    name    VARCHAR(200) NOT NULL
);

CREATE TABLE game_edition_names
(
    language_id     INT REFERENCES languages (id)     NOT NULL,
    game_edition_id INT REFERENCES game_editions (id) NOT NULL,
    name            TEXT                              NOT NULL,
    CONSTRAINT uq_edition_language UNIQUE (language_id, game_edition_id)
);