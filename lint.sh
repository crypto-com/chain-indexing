#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

GOLANGCI_LINT_VERSION=1.27
GOLANGCI_LINT_PATH=${GOLANGCI_LINT_PATH:-golangci-lint}

RET_VALUE=0

check_golangci_lint_version() {
    set +e
    "${GOLANGCI_LINT_PATH}" version 2>&1 | grep "${GOLANGCI_LINT_VERSION}" > /dev/null
    RET_VALUE=$?
    set -e
}

cd "$(dirname "${BASH_SOURCE[0]}")"

echoerr() { echo "$@" 1>&2; }

set +e
command -v "${GOLANGCI_LINT_PATH}" > /dev/null
if [[ $? != 0 ]]; then
    echoerr "Cannot find golangci-lint. Please make sure it is installed."
    exit 1
fi
set -e

check_golangci_lint_version
if [[ "${RET_VALUE}" != 0 ]]; then
    echo "Warning: You are running a different golangci-lint from ours (${GOLANGCI_LINT_VERSION}). The linting result may have some differences."
fi

"${GOLANGCI_LINT_PATH}" run
