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

### Configuration
```
config := bootstrap.Config{
	FileConfig: bootstrap.FileConfig{
		Blockchain: bootstrap.BlockchainConfig{
		    // Bonding denom of the blockchain
			BondingDenom:           "",
		    // Account address prefix of the blockchain
			AccountAddressPrefix:   "",
		    // Account public key prefix of the blockchain
			AccountPubKeyPrefix:    "",
		    // Validator address prefix of the blockchain
			ValidatorAddressPrefix: "",
		    // Validator public key prefix of the blockchain
			ValidatorPubKeyPrefix:  "",
		},
		System: bootstrap.SystemConfig{
		    // "EVENT_STORE", "TENDERMINT_DIRECT", "API_ONLY"
			Mode: "",
		},
		Sync: bootstrap.SyncConfig{
		    // Window size of Sunc process
			WindowSize: 0,
		},
		Tendermint: bootstrap.TendermintConfig{
		    // HTTP address of Tendermint client
			HTTPRPCUrl:           "",
		    // Connection type
			Insecure:             false,
			StrictGenesisParsing: false,
		},
		CosmosApp: bootstrap.CosmosAppConfig{
		    // HTTP address of Cosmos app client
			HTTPRPCUrl: "",
		    // Connection type
			Insecure:   false,
		},
		HTTP: bootstrap.HTTPConfig{
		    // HTTP address to be listened
			ListeningAddress:   "",
		    // Prefix of all routes
			RoutePrefix:        "",
			// Allowed CORS for Origins
			CorsAllowedOrigins: nil,
			// Allowed CORS for Methods
			CorsAllowedMethods: nil,
			// Allowed CORS for Headers
			CorsAllowedHeaders: nil,
		},
		Debug: bootstrap.DebugConfig{
			// Enable pprof server
			PprofEnable:           false,
			// Pprof server address to be listened
			PprofListeningAddress: "",
		},
		Database: bootstrap.DatabaseConfig{
		    // Connection type
			SSL:      false,
		    // Database host
			Host:     "",
		    // Database port
			Port:     0,
		    // Database username
			Username: "",
		    // Database password
			Password: "",
		    // Database name
			Name:     "",
		    // Database schema name
			Schema:   "",
		},
		Postgres: bootstrap.PostgresConfig{
		    // Max connections of Database
			MaxConns:            0,
		    // Min connections of Database
			MinConns:            0,
		    // Max connections life time of Database
			MaxConnLifeTime:     "",
		    // Max connections idle time of Database
			MaxConnIdleTime:     "",
		    // Health check interval of Database
			HealthCheckInterval: "",
		},
		Logger: bootstrap.LoggerConfig{
		    // LOG_LEVEL_DEBUG, LOG_LEVEL_INFO, LOG_LEVEL_ERROR, LOG_LEVEL_PANIC, LOG_DISABLED
			Level: (logger.LogLevel),
			// Enable colered logs
			Color: false,
		},
	    CosmosVersionEnabledHeight: bootstrap.CosmosVersionEnabledHeightConfig{
	            // BLock height from cosmos sdk version v0.42.7
		        V0_42_7: 0,
			},
		},
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
