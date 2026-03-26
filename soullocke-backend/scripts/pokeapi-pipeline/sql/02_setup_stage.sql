DROP SCHEMA IF EXISTS poke_stage CASCADE;
CREATE SCHEMA poke_stage;

CREATE TABLE IF NOT EXISTS poke_stage.version_group
(
    id            bigint PRIMARY KEY,
    name          text   NOT NULL,
    generation_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS poke_stage.version
(
    id               bigint PRIMARY KEY,
    name             text   NOT NULL,
    version_group_id bigint NOT NULL
);