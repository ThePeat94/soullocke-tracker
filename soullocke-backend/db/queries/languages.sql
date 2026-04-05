-- name: FindLanguageByName :one
SELECT id, iso639, iso3166
FROM languages
WHERE name = sqlc.arg(name);

-- name: GetLocalizedNamesForGameEditions :many
SELECT gen.game_edition_id, l.name as language_name, gen.name as edition_name
FROM game_edition_names gen
         INNER JOIN languages l ON gen.language_id = l.id
WHERE gen.language_id = ANY(sqlc.arg(language_ids)::SMALLINT[]);