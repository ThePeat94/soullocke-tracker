#!/usr/bin/env bash
set -euo pipefail
set -a
source .env
set +a

DUMP_FILE_NAME="pokeapi.dump"
ZIPPED_NAME="$DUMP_FILE_NAME.zip"


log() {
  printf '\n[%s] %s\n' "$(date '+%Y-%m-%d %H:%M:%S')" "$1"
}

log "$ZIPPED_NAME"

blank_psql() {
    PGPASSWORD="$DB_PASSWORD" \
    psql \
      -h "$DB_HOST" \
      -p "$DB_PORT" \
      -U "$DB_USERNAME" \
      -d postgres \
      -v ON_ERROR_STOP=1 \
      "$@"
}

scratch_psql() {
  PGPASSWORD="$DB_PASSWORD" \
  psql \
    -h "$DB_HOST" \
    -p "$DB_PORT" \
    -U "$DB_USERNAME" \
    -d "$SCRATCH_DATABASE" \
    -v ON_ERROR_STOP=1 \
    "$@"
}

app_psql() {
    PGPASSWORD="$DB_PASSWORD" \
    psql \
      -h "$DB_HOST" \
      -p "$DB_PORT" \
      -U "$DB_USERNAME" \
      -d "$TARGET_DATABASE" \
      -v ON_ERROR_STOP=1 \
      "$@"
}

restore_into_temp_database() {
  log "restoring poke api data into temporary database"
  PGPASSWORD="$DB_PASSWORD" \
  pg_restore \
   --clean \
   --if-exists \
   --no-owner \
   --no-privileges \
   --single-transaction \
   -d "$SCRATCH_DATABASE" \
   -h "$DB_HOST" \
   -p "$DB_PORT" \
   -U "$DB_USERNAME" \
   "$DUMP_FILE_NAME"
}

download_latest_dump() {
  log "downloading latest dump from pokeapi releases"
  log "requesting latest release"
  LATEST_RELEASE=$(curl -s https://api.github.com/repos/PokeAPI/pokeapi/releases/latest | jq -r '.tag_name')
  log "latest release is $LATEST_RELEASE"
  curl -L -o ./"$ZIPPED_NAME" "https://github.com/PokeAPI/pokeapi/releases/download/$LATEST_RELEASE/pokeapi.dump.zip"
  gunzip -S .zip ./"$ZIPPED_NAME"
}

create_temp_db_and_role() {
  log "creating temporary database"
  blank_psql -f ./sql/01_setup_temp_db.sql
}

teardown_temp_db_and_role() {
  log "tearing down"
  blank_psql -f ./sql/XX_teardown_temp_db.sql
}

cleanup() {
  rm ./"$DUMP_FILE_NAME"
}

copy_from_temp_to_stage() {
  log "start copying data from temp to stage"

  log "setting up stage"
  app_psql -f ./sql/02_setup_stage.sql

  log "copying version groups"
  scratch_psql -Atc "COPY (SELECT id, name, generation_id FROM public.pokemon_v2_versiongroup ORDER BY id) TO STDOUT WITH CSV" \
  | app_psql -c "\COPY poke_stage.version_group (id, name, generation_id) FROM STDIN WITH CSV"

  log "copying versions"
  scratch_psql -Atc "COPY (SELECT id, name, version_group_id FROM public.pokemon_v2_version ORDER BY id) TO STDOUT WITH CSV" \
  | app_psql -c "\COPY poke_stage.version (id, name, version_group_id) FROM STDIN WITH CSV"

  log "transforming data"
  app_psql -f ./sql/03_transform_into_app.sql

  log "tearing down stage"
  app_psql -f ./sql/XX_teardown_stage.sql
}

log "starting etl for pokemon data"
trap 'cleanup; teardown_temp_db_and_role' EXIT
download_latest_dump
create_temp_db_and_role
restore_into_temp_database
copy_from_temp_to_stage
