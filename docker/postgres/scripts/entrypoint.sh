#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

POSTGRES_SSL=${POSTGRES_SSL:-0}
if [[ "${POSTGRES_SSL}" == 1 ]]; then
    ./generate_cert.sh
    mv server.crt /var/lib/postgresql/server.crt
    mv server.key /var/lib/postgresql/server.key
    chown 999:999 /var/lib/postgresql/server.key
    chmod 600 /var/lib/postgresql/server.key

    echo "Running Postgres in SSL mode"
    if [[ "x$1" == "xpostgres" ]]; then
        shift 1
    fi
    ./docker-entrypoint.sh -c ssl=on -c ssl_cert_file=/var/lib/postgresql/server.crt -c ssl_key_file=/var/lib/postgresql/server.key "$@"
else
    ./docker-entrypoint.sh "$@"
fi
