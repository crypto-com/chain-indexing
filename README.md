# Crypto.com Chain Indexing Service

Crypto.com Chain Indexing Service (chain-indexing) is a service to index all publicly available data on Crypto.com chain and persist structured information into storage.

Right now it supports Postgres database and provides RESTful API as query interface.

## 1. Build

### 1.1 Build using Docker (Not working yet)

Using Docker is the easiest way to build chain-indexing. The Docker image contains
`/app/chain-indexing` Compiled chain-indexing binary
`/app/migrate` Program to perform migration

```bash
docker build -o chain-indexing .
```

### 1.2 Build manually

#### Prerequisite

- [Go](https://golang.org/dl/) compiler

```bash
make all
```

Please make sure `$GOPATH` is set and `$GOPATH/bin` includes in `$PATH`. You could include the following code into your `zsh` or `bash` file (`.zshrc` or `.bash_profile`).

```bash
# Add $GOPATH/bin to $PATH
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```

Then `source` the file or open a new terminal session.

## 2. How to Run

### 2.1 Prerequisite

- Crypto.com Chain full node
- Postgres Database

### 2.2 Configuration file

A sample configuration is available under `config/config.sample.toml`.

Copy it, update configuration based on your setup and rename it as `config/config.toml`.

Note: Postgres database password is not available in `config.toml` nor command option. You must provide it as environment variable `DB_PASSWORD` on start.

#### Reminder On Connecting Mainnet

There is a rate limiter on our `mainnet` node. You need to run your indexing server under VPN to pull blocks from `mainnet`.

We have two options at the moment:

- `cdc.gpcloudservice.com`, then choose `gateway` as `HK Office VPN`
- `vpn.mcointernal.com`

### 2.3 Postgres Database

You can have your Postgres setup locally or remotely.

**REMINDER**: I would suggest using our `docker-compose` script to start the DB instance. If you install through `homebrew`, its default setting will need to be adjusted in order to match the indexing server's configuration. 

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

for DB_PASSWORD, never use common word such as "postgres" , "admin", "password", choose at least 16 characters including number, special character, capital letters even in testing environment.
if you don't use strong password, pgmigrate will stop further processing

#### Docker (Not working yet)

```bash
docker run -it \
    --env DB_USERNAME=postgres \
    --env DB_PASSWORD=your_postgresql_password \
    --env DB_HOST=host.docker.internal \
    --env DB_PORT=5432 \
    --env DB_NAME=postgres \
    --env DB_SCHEMA=public \
    chain-indexing /app/migrate -- -verbose up
```

#### Manual Build

```bash
# In your first run, you need to install the dependency `migrate`
./pgmigrate.sh --install-dependency
# Then you should have `migrate` under your `$PATH`
which migrate

# Run the below command to start the migrate
./pgmigrate.sh -- -verbose up
```

### 2.5 Run the Service

#### Docker (Not working yet)

```bash
docker run \
    -v `pwd`/config:/app/config --read-only \
    -p 28857:28857 \
    --env DB_PASSWORD=your_postgresql_password \
    chain-indexing /app/chain-indexing
```

#### Manual build

```bash
env DB_PASSWORD=your_postgresql_password ./chain-indexing
```

## 3. Test

```bash
./test.sh [--install-dependency] [--no-db] [--watch]
```

Providing `--install-dependency` will attempt to install test runner [Ginkgo](https://github.com/onsi/ginkgo) if it is not installed before.

## 4. Lint

### With Local Installed golangci-lint

#### Prerequisite

- [golangci-lint](https://github.com/golangci/golangci-lint)

```bash
./lint.sh
```

### With Docker

```bash
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.33 golangci-lint run -v
```

## 5. Contributing

Please abide by the [Code of Conduct](CODE_OF_CONDUCT.md) in all interactions,
and the [contributing guidelines](CONTRIBUTING.md) when submitting code.

## 6. License

[Apache 2.0](./LICENSE)
