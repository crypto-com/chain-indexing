#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# TODO: Use docker for reproducible environment
DEFAULT_MIGRATION_FOLDER="$(pwd)/migrations"
export TEST_POSTGRES_MIGRATIONS_FOLDER=${TEST_POSTGRES_MIGRATIONS_FOLDER:-${DEFAULT_MIGRATION_FOLDER}}
export TEST_POSTGRES_USERNAME=${TEST_POSTGRES_USERNAME:-postgres}
export TEST_POSTGRES_PASSWORD=${TEST_POSTGRES_PASSWORD:-postgres}
export TEST_POSTGRES_HOST=${TEST_POSTGRES_HOST:-localhost}
export TEST_POSTGRES_PORT=${TEST_POSTGRES_PORT:-5433}
export TEST_POSTGRES_DATABASE=${TEST_POSTGRES_DATABASE:-postgres_test}
export TEST_POSTGRES_SCHEMA=${TEST_POSTGRES_SCHEMA:-public}

DOCKER_COMPOSE_PROJECT="chain_indexing_test"

FN_UNKNOWN="unknown"
FN_BOOTSTRAP="bootstrap"
FN_GENERATE="generate"
FN_SETUP="setup"
FN_TEARDOWN="teardown"
FN_TEST="test"

# main
RET_VALUE=0

cd "$(dirname "${BASH_SOURCE[0]}")"

POSTGRES_DRIVER_URL="postgres://${TEST_POSTGRES_USERNAME}:${TEST_POSTGRES_PASSWORD}@${TEST_POSTGRES_HOST}:${TEST_POSTGRES_PORT}/${TEST_POSTGRES_DATABASE}?options=--search_path%3d=${TEST_POSTGRES_SCHEMA}"

echoerr() { echo "$@" 1>&2; }

check_ginkgo() {
    set +e
    command -v ginkgo > /dev/null
    RET_VALUE=$?
    set -e
}

trap handle_ctrlc INT
handle_ctrlc() {
    echo "Captured Ctrl-C. Stopping test environment ..."
    teardown

    echo
    echo "Test Interrupted"
    exit 1
}

teardown() {
    if [[ "${TEST_DB}" == 1 ]]; then
        docker-compose -f ./docker/docker-compose.test.yml -p "${DOCKER_COMPOSE_PROJECT}" down
    fi
}

setup() {
    if [[ "${TEST_DB}" == 1 ]]; then
        docker-compose -f ./docker/docker-compose.test.yml -p "${DOCKER_COMPOSE_PROJECT}" up -d
        wait_postgres_ready
    fi
}

wait_postgres_ready() {
    CONTAINER_ID=$(docker ps -aqf "name=${DOCKER_COMPOSE_PROJECT}_postgres")
    
    while :
    do
        is_postgres_ready "${CONTAINER_ID}"
        if [[ "${RET_VALUE}" == 0 ]]; then
            echo "Postgres is ready to accept connections"
            return
        fi
        docker ps
        echo "Waiting for Postgres container(${CONTAINER_ID}) to be ready ..."
        sleep 1
    done
}

is_postgres_ready() {
    set +e
    docker exec "$1" pg_isready > /dev/null
    RET_VALUE=$?
    set -e
}

run_test() {
    set +e
    ginkgo -r
    RET_VALUE=$?
    if [[ "${RET_VALUE}" != 0 ]]; then
      return
    fi
    # TODO: Migrate coin test cases to Ginkgo
    go test ./usecase/coin/...
    RET_VALUE=$?
    set -e
}

run_test_watch() {
    set +e
    ginkgo watch -r
    RET_VALUE=$?
    set -e
}

run_bootstrap() {
    pushd .
    cd $1
    ginkgo bootstrap
    RET_VALUE=$?
    popd
}

run_generate() {
    pushd .
    cd $(dirname $1)
    ginkgo generate $(basename $1)
    RET_VALUE=$?
    popd
}

INSTALL_DEPENDENCY=0
WATCH_MODE=0
TEST_DB=1
FN=FN_UNKNOWN
while [[ $# > 0 ]]; do
    case "$1" in
        bootstrap)
            FN=FN_BOOTSTRAP
            shift 1
        ;;
        generate)
            FN=FN_GENERATE
            shift 1
        ;;
        setup)
            FN=FN_SETUP
            shift 1
        ;;
        teardown)
            FN=FN_TEARDOWN
            shift 1
        ;;
        run)
            FN=FN_TEST
            shift 1
        ;;
        --install-dependency)
            INSTALL_DEPENDENCY=1
            shift 1
        ;;
        *)
            break
        ;;
    esac
done

check_ginkgo
if [[ "${RET_VALUE}" != 0 ]]; then
    if [[ "${INSTALL_DEPENDENCY}" = 0 ]]; then
        echoerr "Cannot find ginkgo. Provide \`--install-dependency\` to attempt to install it."
        exit 1
    fi
    
    go get -u github.com/onsi/ginkgo/ginkgo
    check_ginkgo
    if [[ "${RET_VALUE}" != 0 ]]; then
        echoerr "Failed to install ginkgo. Please install it manually."
        exit 1
    fi
fi

if [[ "${FN}" == FN_BOOTSTRAP ]]; then
    run_bootstrap $@
    exit "${RET_VALUE}"
elif [[ "${FN}" == FN_GENERATE ]]; then
    run_generate $@
    exit "${RET_VALUE}"
elif [[ "${FN}" == FN_SETUP ]]; then
    echo "Stopping any previous test environment ..."
    teardown
    setup
elif [[ "${FN}" == FN_TEARDOWN ]]; then
    echo "Stopping any previous test environment ..."
    teardown
elif [[ "${FN}" == FN_TEST ]]; then
    while [[ $# > 0 ]]; do
        case "$1" in
            --no-db)
                TEST_DB=0
                shift 1
            ;;
            --no-ssl)
                POSTGRES_DRIVER_URL="${POSTGRES_DRIVER_URL}&sslmode=disable"
                shift 1
            ;;
            --watch)
                WATCH_MODE=1
                shift 1
            ;;
            *)
                break
            ;;
        esac
    done

    echo "Stopping any previous test environment ..."
    teardown
    setup

    export TEST_DATABASE_URL="${POSTGRES_DRIVER_URL}"
    echo
    if [[ "${WATCH_MODE}" == 1 ]]; then
        run_test_watch
    else
        run_test
    fi
    TEST_RESULT="${RET_VALUE}"
    echo

    echo "Test finished. Stopping test envionrment ..."
    teardown

    echo
    if [[ "${TEST_RESULT}" != 0 ]]; then
        echo "Test Failed. Check the trace above"
    else
        echo "Test Passed"
    fi
    exit "${TEST_RESULT}"
else
    echo "Unknown or missing command"
    exit 1
fi