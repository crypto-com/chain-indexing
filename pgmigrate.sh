#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# TODO: Use docker for reproducible environment
# Configurations
DB_MIGRATIONS_FOLDER=${DB_MIGRATIONS_FOLDER:-./migrations}
DB_PROTOCOL=${DB_PROTOCOL:-postgres}
DB_USERNAME=${DB_USERNAME:-postgres}
DB_PASSWORD=${DB_PASSWORD:-postgres}
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_NAME=${DB_NAME:-postgres}
DB_SCHEMA=${DB_SCHEMA:-public}

# main
RET_VALUE=0

cd "$(dirname "${BASH_SOURCE[0]}")"

DB_DRIVER_URL="${DB_PROTOCOL}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?options=--search_path=${DB_SCHEMA}"

echoerr() { echo "$@" 1>&2; }

check_migrate() {
    set +e
    command -v migrate > /dev/null
    RET_VALUE=$?
    set -e
}

run_migrate() {
    set -x
    migrate "$@"
    exit $?
}

INSTALL_DEPENDENCY=0
while [[ $# > 0 ]]; do
    case "$1" in
        --install-dependency)
            INSTALL_DEPENDENCY=1
            shift 1
        ;;
        --no-ssl)
            DB_DRIVER_URL="${DB_DRIVER_URL}&sslmode=disable"
            shift 1
        ;;
        --)
            shift 1
            break
        ;;
        *)
            echoerr "Unknown argument: $1"
            exit 1
        ;;
    esac
done

check_migrate
if [[ "${RET_VALUE}" != 0 ]]; then
    if [[ "${INSTALL_DEPENDENCY}" = 0 ]]; then
        echoerr "Cannot find migrate. Provide \`--install-dependency\` to attempt to install it."
        exit 1
    fi
    
    go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
    check_migrate
    if [[ "${RET_VALUE}" != 0 ]]; then
        echoerr "Failed to install migrate. Please install it manually."
        exit 1
    fi
fi

if [[ $# == 0 ]]; then
    echoerr "Missing command"
    exit 1
fi
ARGS="-verbose"
while [[ $# > 0 ]]; do
    case "$1" in
        create)
            shift 1
            run_migrate "${ARGS}" create --dir "${DB_MIGRATIONS_FOLDER}" --ext sql "$@"
        ;;
        goto | up | down | drop | force)
            run_migrate -source "file://${DB_MIGRATIONS_FOLDER}" -database "${DB_DRIVER_URL}" "${ARGS}" "$@"
        ;;
        version)
            run_migrate version
        ;;
        help)
            run_migrate "${ARGS}" help
        ;;
        -verbose)
            shift 1
        ;;
        *)
            ARGS="${ARGS} $1"
            shift 1
        ;;
    esac
done
echoerr "Unknown migrate command: ${ARGS}"
