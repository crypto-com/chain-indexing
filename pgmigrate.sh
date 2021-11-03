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

GOLANG_MIGRATE_VERSION="v4.14.1"
check_migrate() {
    set +e
    command -v migrate > /dev/null
    if [[ $? != 0 ]]; then
        RET_VALUE="NOT_FOUND"
        set -e
        return
    fi
    VERSION=$(migrate --version 2>&1)
    if [[ "${VERSION}" != "${GOLANG_MIGRATE_VERSION}" ]]; then
        if [[ $# == 1 && "$1" == "--verbose" ]]; then
            echo "WARNING: System is running migrate version ${VERSION}. For best compatibility, we recommended using ${GOLANG_MIGRATE_VERSION}. Consider provide \`--install-dependency\` to try to install this version."
        fi
        RET_VALUE="VERSION_MISMATCH"
        set -e
        return
    fi
    RET_VALUE="OK"
    set -e
}

is_golang_1_16_and_above() {
    IFS=$' '
    GOVERSION="$(go version | { read _ _ v _; echo ${v#go}; })"
    IFS=$'\n\t'
    echo ${GOVERSION}
    MAJOR="$(echo ${GOVERSION} | cut -d. -f1)"
    MINOR="$(echo ${GOVERSION} | cut -d. -f2)"
    REVISION="$(echo ${GOVERSION} | cut -d. -f3)"
    if [[ "${MAJOR}" == "" ]]; then
        echoerr "Golang version not found. Make sure it is installed."
        exit 1
    fi

    if [[ $MAJOR != 1 ]]; then
        echoerr "Only Golang v1 is supported. Found v${GOVERSION}"
        exit 1
    fi

    if [[ $MINOR -ge 16 ]]; then
        RET_VALUE=1
    else
        RET_VALUE=0
    fi
}

# https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#versioned
install_golang_migrate_golang_1_15_and_below() {
    pushd .
    set -x
    # https://github.com/golang-migrate/migrate/pull/257#issuecomment-705249902
    GO111MODULE=off go get -tags 'postgres' -u -d github.com/golang-migrate/migrate/cmd/migrate/
    cd "${GOPATH}/src/github.com/golang-migrate/migrate/cmd/migrate"
    git checkout ${GOLANG_MIGRATE_VERSION}
    go build -tags 'postgres' -ldflags="-X main.Version=$(git describe --tags)" -o "${GOPATH}/bin/migrate" "${GOPATH}/src/github.com/golang-migrate/migrate/cmd/migrate"
    set +x
    popd
}

install_golang_migrate_universal() {
    install_golang_migrate_golang_1_15_and_below
}

install_golang_migrate_golang_1_16_and_above() {
    # This is the recommended way to install golang-migrate. But the binary installed using this method will always
    # return `dev` when running `migrate --version` and will confuse the `migrate` version check.
    set -x
    go install -tags 'postgres' "github.com/golang-migrate/migrate/v4/cmd/migrate@${GOLANG_MIGRATE_VERSION}"
    set +x
}

run_migrate() {
    migrate "$@"
    exit $?
}

INSTALL_DEPENDENCY=0
INSECURE=0
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

check_migrate --verbose
if [[ "${RET_VALUE}" != "OK" ]]; then
    if [[ "${INSTALL_DEPENDENCY}" == 0 ]]; then
        if [[ "${RET_VALUE}" == "NOT_FOUND" ]]; then
            echoerr "Cannot find migrate. Provide \`--install-dependency\` to try to install it."
            exit 1
        fi
    else
        install_golang_migrate_universal

        check_migrate
        if [[ "${RET_VALUE}" == "NOT_FOUND" ]]; then
            echoerr "Failed to install migrate. Please install it manually."
            exit 1
        fi
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
            run_migrate --version
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
