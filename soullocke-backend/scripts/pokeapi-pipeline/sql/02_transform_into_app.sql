BEGIN;

DO
$$
    DECLARE
        total_count    int;
        inserted_count int;
    BEGIN
        SELECT count(*)
        INTO total_count
        FROM poke_stage.version v
                 JOIN poke_stage.version_group vg ON vg.id = v.version_group_id;

        WITH inserted AS (
            INSERT INTO game_editions (id, name, generation)
                SELECT v.identifier, initcap(replace(v.identifier, '-', ' ')), vg.generation_id
                FROM poke_stage.version v
                         JOIN poke_stage.version_group vg ON vg.id = v.version_group_id
                ON CONFLICT DO NOTHING
                RETURNING id)
        SELECT count(*)
        INTO inserted_count
        FROM inserted;

        RAISE NOTICE '% of % rows inserted (% skipped)', inserted_count, total_count, total_count - inserted_count;
    END
$$;

COMMIT;