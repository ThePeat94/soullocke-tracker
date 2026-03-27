#!/usr/bin/env bash
set -euo pipefail
set -a
source .env
set +a

REPO_BASE_URL="https://raw.githubusercontent.com/PokeAPI/pokeapi"
REF_COMMIT_SHA="8711df8f5216c2ea5698a779bedd4ef7e2166059"


get_file_url() {
  echo "$REPO_BASE_URL/$REF_COMMIT_SHA/data/v2/csv/$1.csv"
}

log() {
  printf '\n[%s] %s\n' "$(date '+%Y-%m-%d %H:%M:%S')" "$1"
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

download_latest_dumps() {
  log "downloading csv files from github repo"
  mkdir -p ./tmp
  curl -L -o ./tmp/versions.csv "$(get_file_url "versions")"
  curl -L -o ./tmp/version_groups.csv "$(get_file_url "version_groups")"
}

cleanup() {
  log "cleaning up"
  app_psql -f ./sql/03_teardown_stage.sql 2>/dev/null || true
  rm -rf ./tmp
}

copy_from_temp_to_stage() {
  log "start copying data from temp to stage"

  log "setting up stage"
  app_psql -f ./sql/01_setup_stage.sql

  log "copying version groups"
  app_psql -c "\COPY poke_stage.version_group (id, identifier, generation_id, \"order\") FROM './tmp/version_groups.csv' WITH (FORMAT csv, HEADER true)"

  log "copying versions"
  app_psql -c "\COPY poke_stage.version (id, version_group_id, identifier) FROM './tmp/versions.csv' WITH (FORMAT csv, HEADER true)"

  log "transforming data"
  app_psql -f ./sql/02_transform_into_app.sql
}

log "starting etl for pokemon data"
trap 'cleanup' EXIT
download_latest_dumps
copy_from_temp_to_stage
