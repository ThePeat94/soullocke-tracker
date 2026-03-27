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
