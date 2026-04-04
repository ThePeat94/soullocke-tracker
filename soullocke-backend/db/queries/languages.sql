-- name: FindLanguageByName :one
SELECT id, iso639, iso3166
FROM languages
WHERE name = sqlc.arg(name);

-- name: GetLocalizedNamesForGameEdition :many
SELECT l.name as language_name, gen.name as edition_name
FROM game_edition_names gen
INNER JOIN public.languages l on gen.language_id = l.id
WHERE gen.game_edition_id = sqlc.arg(game_edition_id)
  AND gen.language_id = ANY (sqlc.arg(language_ids)::SMALLINT[]);