# Crypto.com Chain Indexing Service

Crypto.com Chain Indexing Service (chainindex) is a service to index all publicly available data on Crypto.com chain and persist structured information into storage.

Right now it supports Postgres database and provides RESTful API as query interface.

## 1. Build

### 1.1 Build using Docker

Using Docker is the easiest way to build chainindex. The Docker image contains
`/app/chainindex` Compiled chainindex binary
`/app/migrate` Program to perform migration

```bash
docker build -o chainindex .
```

### 1.2 Build manually

#### Prerequisite

- [Go](https://golang.org/dl/) compiler
- [Rust](https://rustup.rs/) compiler

```bash
make all
```

## 2. How to Run

### 2.1 Prerequisite

- Crypto.com Chain full node
- Postgres Database

### 2.2 Configuration file

A sample configuration is available under `config/config.sample.toml`.

Copy it, update configuration based on your setup and rename it as `config/config.toml`.

Note: Postgres database password is not available in `config.toml` nor command option. You must provide it as environment variable `DB_PASSWORD` on start.

### 2.3 Postgres Database

You can have your Postgres setup locally or remotely.

#### Run Postgres with Docker

**WARNING**: The docker files available under `docker/` is intended only for development and testing purpose only. Never use them in production setup.

For a local test run. A docker-compose file with Postgres database and PgAdmin console is available via running:

```bash
docker-compose --file docker/docker-compose.development.yml up -d
```

This will start the following docker instances on your local network when you use default credentials:
| Docker image | Port | Username | Password | Other Config | Mounted volume |
| --- | --- | --- | --- | --- | --- |
| Postgres | 5432 | postgres | postgres | Database Name = postgres; SSL = true | ./pgdata-dev |
| PgAdmin | 8080 | pgadmin@localhost | pgadmin | N/A | N/A |

### 2.4 Execute Database Migration

#### Docker

```bash
docker run -it \
    --env DB_USERNAME=postgres \
    --env DB_PASSWORD=postgres \
    --env DB_HOST=host.docker.internal \
    --env DB_PORT=5432 \
    --env DB_NAME=postgres \
    --env DB_SCHEMA=public \
    chainindex /app/migrate -- -verbose up
```

#### Manual Build

```bash
./migrate.sh -- -verbose up
```

### 2.5 Run the Service

#### Docker

```bash
docker run \
    -v `pwd`/config:/app/config --read-only \
    -p 28857:28857 \
    --env DB_PASSWORD=postgres \
    chainindex /app/chainindex
```

#### Manual build

```bash
env DB_PASSWORD=postgres ./chainindex
```

## 3. Test

```bash
./test.sh [--install-dependency] [--no-db] [--watch]
```

Providing `--install-dependency` will attempt to install test runner [Ginkgo](https://github.com/onsi/ginkgo) if it is not installed before.

## 4. Lint

#### Prerequisite

- [golangci-lint](https://github.com/golangci/golangci-lint)

```bash
./lint.sh
```

## 5. Contributing

Please abide by the [Code of Conduct](CODE_OF_CONDUCT.md) in all interactions,
and the [contributing guidelines](CONTRIBUTING.md) when submitting code.

## 6. License

[Apache 2.0](./LICENSE)