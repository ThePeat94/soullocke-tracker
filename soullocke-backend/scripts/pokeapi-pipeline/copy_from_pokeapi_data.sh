#!/usr/bin/env bash

# This script describes an ETL workflow to fill the database with relevant data

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

set -a
source .env
set +a

REPO_BASE_URL="https://raw.githubusercontent.com/PokeAPI/pokeapi"
# PokeAPI repo commit to pull CSV data from — update this when new Pokemon generations release
REF_COMMIT_SHA="8711df8f5216c2ea5698a779bedd4ef7e2166059"

# format: "table_name:col1,col2,col3"
STAGE_TABLES=(
  "version_groups:id,identifier,generation_id,\"order\""
  "versions:id,version_group_id,identifier"
  "languages:id,iso639,iso3166,identifier,official,\"order\""
  "version_names:version_id,local_language_id,name"
)

get_file_url() {
  echo "$REPO_BASE_URL/$REF_COMMIT_SHA/data/v2/csv/$1.csv"
}

log() {
  printf '[%s] %s\n' "$(date '+%Y-%m-%d %H:%M:%S')" "$1"
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
   for entry in "${STAGE_TABLES[@]}"; do
    FILE=${entry%%:*}
    log "downloading $FILE"
    FILE_URL=$(get_file_url "$FILE")
    if ! curl --fail -L -o "./tmp/$FILE.csv" "$FILE_URL"; then
      log "ERROR: failed to download $FILE from $FILE_URL"
      exit 1
    fi
  done
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

  for entry in "${STAGE_TABLES[@]}"; do
    table="${entry%%:*}"
    columns="${entry#*:}"
    log "copying $table into stage"
    app_psql -c "\COPY poke_stage.$table ($columns) FROM './tmp/${table}.csv' WITH (FORMAT csv, HEADER true)"
  done

  log "transforming data"
  app_psql -f ./sql/02_transform_into_app.sql
}

log "starting etl for pokemon data"
trap 'cleanup' EXIT
download_latest_dumps
copy_from_temp_to_stage
