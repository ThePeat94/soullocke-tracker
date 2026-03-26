#!/usr/bin/env bash
set -euo pipefail
set -a
source .env
set -a

DUMP_FILE_NAME="pokeapi.dump"
ZIPPED_NAME="$DUMP_FILE_NAME.zip"


log() {
  printf '\n[%s] %s\n' "$(date '+%Y-%m-%d %H:%M:%S')" "$1"
}

log "$ZIPPED_NAME"

init_psql() {
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

cleanup() {
  rm ./"$DUMP_FILE_NAME"
}

log "starting etl for pokemon data"
download_latest_dump
log "creating temporary database"
init_psql -f ./sql/01_setup_temp_db.sql
restore_into_temp_database
log "tearing down"
init_psql -f ./sql/XX_teardown_temp_db.sql
cleanup