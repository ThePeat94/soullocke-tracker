BEGIN;

-- insert game editions
DO
$$
    DECLARE
        total_count    int;
        inserted_count int;
    BEGIN
        SELECT count(*)
        INTO total_count
        FROM poke_stage.versions;

        WITH inserted AS (
            INSERT INTO game_editions (id, code_name, generation, fallback_name)
                SELECT v.id, v.identifier, vg.generation_id, initcap(replace(v.identifier, '-', ' '))
                FROM poke_stage.versions v
                         JOIN poke_stage.version_groups vg ON vg.id = v.version_group_id
                ON CONFLICT DO NOTHING
                RETURNING id)
        SELECT count(*)
        INTO inserted_count
        FROM inserted;

        RAISE NOTICE 'game editions: % of % rows inserted (% skipped)', inserted_count, total_count, total_count - inserted_count;
    END
$$;

-- insert languages
DO
$$
    DECLARE
        total_count    int;
        inserted_count int;
    BEGIN
        SELECT count(*)
        INTO total_count
        FROM poke_stage.languages;

        WITH inserted AS (
            INSERT INTO languages (id, iso639, iso3166, name)
                SELECT l.id, l.iso639, l.iso3166, l.identifier
                FROM poke_stage.languages l
                ON CONFLICT DO NOTHING
                RETURNING id)
        SELECT count(*)
        INTO inserted_count
        FROM inserted;

        RAISE NOTICE 'languages: % of % rows inserted (% skipped)', inserted_count, total_count, total_count - inserted_count;
    END
$$;


-- insert game edition names
DO
$$
    DECLARE
        total_count    int;
        inserted_count int;
    BEGIN
        SELECT count(*)
        INTO total_count
        FROM poke_stage.version_names;

        WITH inserted AS (
            INSERT INTO game_edition_names (language_id, game_edition_id, name)
                SELECT n.local_language_id, n.version_id, n.name
                FROM poke_stage.version_names n
                ON CONFLICT DO NOTHING
                RETURNING *)
        SELECT count(*)
        INTO inserted_count
        FROM inserted;

        RAISE NOTICE 'version names: % of % rows inserted (% skipped)', inserted_count, total_count, total_count - inserted_count;
    END
$$;

COMMIT;
