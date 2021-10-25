# Crypto.com Chain Indexing Service

Crypto.com Chain Indexing Service (chain-indexing) is a service to index all publicly available data on Crypto.com chain and persist structured information into storage.

Right now it supports Postgres database and provides RESTful API as query interface.

## 1. Usage

```go
package main

import (
	"os"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/bootstrap"
	"github.com/crypto-com/chain-indexing/infrastructure"
	"github.com/crypto-com/chain-indexing/entity/projection"
)

func main() {
	// Init configurations...
	logger := infrastructure.NewZerologLogger(os.Stdout)
	fileConfig := bootstrap.FileConfig{}
	config := bootstrap.Config{
		FileConfig: fileConfig,
	}

	// Init indexing app
	app := bootstrap.NewApp(logger, &config)
	app.InitIndexService(
		initProjections(logger, &config),
		initCronJobs(logger, &config),
	)
	app.InitHTTPAPIServer(initHTTPAPIHandlers(logger, &config))

	// Run indexing app
	app.Run()
}

func initProjections(
	logger applogger.Logger,
	config *bootstrap.Config,
) []projection.Projection {
    // append your Projections
}


func initCronJobs(
	logger applogger.Logger,
	config *bootstrap.Config,
) []projection.CronJob {
	// append your CronJobs
}

func initHTTPAPIHandlers(
	logger applogger.Logger,
	config *bootstrap.Config,
) []bootstrap.Handler {
	// append your Handlers
}
```


## 2. Test

```bash
./test.sh [--install-dependency] [--no-db] [--watch]
```

Providing `--install-dependency` will attempt to install test runner [Ginkgo](https://github.com/onsi/ginkgo) if it is not installed before.

## 3. Lint

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

## 4. Contributing

Please abide by the [Code of Conduct](CODE_OF_CONDUCT.md) in all interactions,
and the [contributing guidelines](CONTRIBUTING.md) when submitting code.

## 5. License

[Apache 2.0](./LICENSE)
