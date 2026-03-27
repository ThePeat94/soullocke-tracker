DROP SCHEMA IF EXISTS poke_stage CASCADE;
CREATE SCHEMA poke_stage;

CREATE TABLE IF NOT EXISTS poke_stage.version_group
(
    id            BIGINT PRIMARY KEY,
    identifier    TEXT   NOT NULL,
    generation_id BIGINT NOT NULL,
    "order"       BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS poke_stage.version
(
    id               BIGINT PRIMARY KEY,
    version_group_id BIGINT NOT NULL,
    identifier       TEXT   NOT NULL
);