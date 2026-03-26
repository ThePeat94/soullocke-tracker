BEGIN;

INSERT INTO game_editions (id, name, generation)
SELECT v.name, initcap(replace(v.name, '-', ' ')), vg.generation_id
FROM poke_stage.version v
JOIN poke_stage.version_group vg ON vg.id = v.version_group_id
ON CONFLICT DO NOTHING;

COMMIT;