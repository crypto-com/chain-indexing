#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

cd "$(dirname "${BASH_SOURCE[0]}")"

# WARNING: This scripts generates a weak self-signed certifcate for development
# purpose only. Never use it on production.

echo "Generating self-signed certifcate for Postgres"

openssl req -new -text -passout pass:postgres -subj /CN=localhost -out server.req -keyout privkey.pem
openssl rsa -in privkey.pem -passin pass:postgres -out server.key
openssl req -x509 -in server.req -text -key server.key -out server.crt
