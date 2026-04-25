DROP SCHEMA IF EXISTS poke_stage CASCADE;
CREATE SCHEMA poke_stage;

CREATE TABLE poke_stage.version_groups
(
    id            BIGINT PRIMARY KEY,
    identifier    TEXT   NOT NULL,
    generation_id BIGINT NOT NULL,
    "order"       BIGINT NOT NULL
);

CREATE TABLE poke_stage.versions
(
    id               BIGINT PRIMARY KEY,
    version_group_id BIGINT NOT NULL,
    identifier       TEXT   NOT NULL
);

CREATE TABLE poke_stage.languages
(
    id         BIGINT PRIMARY KEY,
    iso639     VARCHAR(10)  NOT NULL, -- language code
    iso3166    VARCHAR(2)   NOT NULL, -- country code
    identifier varchar(200) NOT NULL,
    official   BOOLEAN      NOT NULL,
    "order"    INTEGER
);

CREATE TABLE poke_stage.version_names
(
    version_id        INTEGER REFERENCES poke_stage.versions (id),
    local_language_id INTEGER REFERENCES poke_stage.languages (id),
    name              TEXT
);