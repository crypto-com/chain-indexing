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
	
	// filling fileConfig
	// ...
	
	config := bootstrap.Config{
		FileConfig: fileConfig,
	}

	// Init indexing app
	app := bootstrap.NewApp(logger, &config)
	app.InitIndexService(
		initProjections(logger, &config),
		initCronJobs(logger, &config),
	)
	app.InitHTTPAPIServer(initRouteRegistry(logger, &config))

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

func initRouteRegistry(
	logger applogger.Logger,
	config *bootstrap.Config,
) bootstrap.RouteRegistry {
	// append your Routes
}
```

### Configuration
```go
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
        GithubAPI: bootstrap.GithubAPIConfig{
        	// Username of your git hub api account
            Username:           "username",
			// Token of your git hub api where at least have public repo access right
            Token:              "token",
            // Specific branch, tag or commit. Leave it empty if always using the latest master
            MigrationRepoRef:   "ref",
        },
        Prometheus: bootstrap.PrometheusConfig{
            Enable:     true,
            ExportPath: "/metrics",
            Port:       "9090",
        },
}
```

### Initial projections
```go
package main

import (
	"github.com/ettle/strcase"

	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/bootstrap"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	cosmosapp_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/cosmosapp"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	github_migrationhelper "github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper/github"
	"github.com/crypto-com/chain-indexing/projection/account"
	"github.com/crypto-com/chain-indexing/projection/account_transaction"
)

func initProjections(
    logger applogger.Logger,
    rdbConn rdb.Conn,
    config *bootstrap.Config,
    customConfig *CustomConfig,
) (projections []projection_entity.Projection) {
    // Skip if API_ONLY is on
    if config.System.Mode == bootstrap.SYSTEM_MODE_API_ONLY {
        return projections
    }

    connString := rdbConn.(*pg.PgxConn).ConnString()
    
    githubMigrationHelperConfig := github_migrationhelper.Config{
        GithubAPIUser:    config.GithubAPI.Username,
        GithubAPIToken:   config.GithubAPI.Token,
        MigrationRepoRef: config.GithubAPI.MigrationRepoRef,
        ConnString:       connString,
    }

    var cosmosAppClient cosmosapp.Client
    if config.CosmosApp.Insecure {
        cosmosAppClient = cosmosapp_infrastructure.NewInsecureHTTPClient(
            config.CosmosApp.HTTPRPCUrl, config.Blockchain.BondingDenom,
        )
    } else {
        cosmosAppClient = cosmosapp_infrastructure.NewHTTPClient(
            config.CosmosApp.HTTPRPCUrl, config.Blockchain.BondingDenom,
        )
    }

    sourceURL := github_migrationhelper.GenerateDefaultSourceURL("Account", githubMigrationHelperConfig)
    databaseURL := migrationhelper.GenerateDefaultDatabaseURL("Account", connString)
    migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)
    
    // Append `Account` projection
    projections = append(account.NewAccount(logger, rdbConn, config.Blockchain.AccountAddressPrefix, cosmosAppClient, migrationHelper). projections)

    sourceURL = github_migrationhelper.GenerateDefaultSourceURL("AccountTransaction", githubMigrationHelperConfig)
    databaseURL = migrationhelper.GenerateDefaultDatabaseURL("AccountTransaction", connString)
    migrationHelper = github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

    projections = append(account_transaction.NewAccountTransaction(logger, rdbConn, config.Blockchain.AccountAddressPrefix, migrationHelper), projections)

    for _, projection := range projections {
        if onInitErr := projection.OnInit(); onInitErr != nil {
            logger.Errorf(
			    "error initializing projection %s: %v",
			    projection.Id(), onInitErr,
            )
        }
    }

    return projections
}
```

#### Custom projection

```go
package example

import (
	"fmt"
	"log"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	example_view "your_view_packge"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

type AdditionalExampleProjection struct {
	*rdbprojectionbase.Base

	rdbConn      rdb.Conn
	logger       applogger.Logger
}

func NewAdditionalProjection(
	logger applogger.Logger,
	rdbConn rdb.Conn,
) *AdditionalExampleProjection {
	return &AdditionalExampleProjection{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "Example"),
		rdbConn,
		logger,
	}
}

var (
	NewExamplesView              = example_view.NewExamplesView
	UpdateLastHandledEventHeight = (*AdditionalExampleProjection).UpdateLastHandledEventHeight
)

func (_ *AdditionalExampleProjection) GetEventsToListen() []string {
	return event_usecase.MSG_EVENTS
}

func (projection *AdditionalExampleProjection) OnInit() error {
	return nil
}

func (projection *AdditionalExampleProjection) HandleEvents(height int64, events []event_entity.Event) error {
	rdbTx, err := projection.rdbConn.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}

	committed := false
	defer func() {
		if !committed {
			_ = rdbTx.Rollback()
		}
	}()

	rdbTxHandle := rdbTx.ToHandle()

	examplesView := NewExamplesView(rdbTxHandle)

	for _, event := range events {
		log.Println(event)
		if typedEvent, ok := event.(*event_usecase.MsgSend); ok {
			row := &example_view.ExampleRow{
				Address: typedEvent.ToAddress,
				Balance: typedEvent.Amount,
			}
			if handleErr := projection.handleSomeEvent(examplesView, row); handleErr != nil {
				return fmt.Errorf("error handling MsgSend: %v", handleErr)
			}
		}
	}

	if err = UpdateLastHandledEventHeight(projection, rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err = rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true

	return nil
}

func (projection *AdditionalExampleProjection) handleSomeEvent(examplesView example_view.Examples, row *example_view.ExampleRow) error {
	return examplesView.Insert(row)
}
```

```go
package view

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/external/json"

	"github.com/crypto-com/chain-indexing/usecase/coin"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	_ "github.com/crypto-com/chain-indexing/test/factory"
)

type Examples interface {
	Insert(*ExampleRow) error
}

type ExamplesView struct {
	rdb *rdb.Handle
}

func NewExamplesView(handle *rdb.Handle) Examples {
	return &ExamplesView{
		handle,
	}
}

func (exampleView *ExamplesView) Insert(example *ExampleRow) error {
	sql, sqlArgs, err := exampleView.rdb.StmtBuilder.
		Insert(
			"view_examples",
		).
		Columns(
			"address",
			"balance",
		).
		Values(
			example.Address,
			json.MustMarshalToString(example.Balance),
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building examples insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := exampleView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting example into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting example into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

type ExampleRow struct {
	Address string     `json:"address"`
	Balance coin.Coins `json:"balance"`
}
```

Append custom projection
```go
func initProjections(
    logger applogger.Logger,
    rdbConn rdb.Conn,
    config *bootstrap.Config,
    customConfig *CustomConfig,
) (projections []projection.Projection) {
    // ...

    githubMigrationHelperConfigForCustomProjection := github_migrationhelper.Config{
        GithubAPIUser:    config.GithubAPI.Username,
        GithubAPIToken:   config.GithubAPI.Token,
        MigrationRepoRef: customConfig.ServerGithubAPI.MigrationRepoRef,
        ConnString:       connString,
    }

    sourceURL := generateGithubMigrationSrouceURLForCustomProjection("Example", githubMigrationHelperConfigForCustomProjection)
    databaseURL := migrationhelper.GenerateDefaultDatabaseURL("Example", connString)
    migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

    projections = append(account_transaction.NewAccountTransaction(params.Logger, rdbConn, migrationHelper), projections)

    return projections
}
```

### Initial CronJobs
```go
package main

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/bootstrap"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	github_migrationhelper "github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper/github"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/bridge_activity_matcher"
)

func initCronJobs(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config *bootstrap.Config,
	customConfig *CustomConfig,
) (crons []projection_entity.CronJob) {
    // Skip if API_ONLY is on
    if config.System.Mode == bootstrap.SYSTEM_MODE_API_ONLY {
        return crons
    }

    connString := rdbConn.(*pg.PgxConn).ConnString()
    
    sourceURL := github_migrationhelper.GenerateSourceURL(
        github_migrationhelper.MIGRATION_GITHUB_URL_FORMAT,
        config.GithubAPI.Username,
        config.GithubAPI.Token,
        bridge_activity_matcher.MIGRATION_DIRECOTRY,
        config.GithubAPI.MigrationRepoRef,
    )
    databaseURL := migrationhelper.GenerateDefaultDatabaseURL("BridgeActivityMatcher", connString)
    migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)
    
    // Append `BridgeActivityMatcher` cron
    crons = append(bridge_activity_matcher.New(logger, rdbConn, migrationHelper). crons)

    for _, cron := range crons {
        if onInitErr := cron.OnInit(); onInitErr != nil {
            logger.Errorf(
                "error initializing cronjob %s: %v",
			    cron.Id(), onInitErr,
            )
        }
    }
    return crons
}
```

### Initial route
```go
package routes

import (
    "github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
    "github.com/crypto-com/chain-indexing/appinterface/rdb"
    "github.com/crypto-com/chain-indexing/appinterface/tendermint"
    "github.com/crypto-com/chain-indexing/bootstrap"
    applogger "github.com/crypto-com/chain-indexing/external/logger"
    cosmosapp_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/cosmosapp"
    httpapi_handlers "github.com/crypto-com/chain-indexing/infrastructure/httpapi/handlers"
    tendermint_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/tendermint"
)

func InitRouteRegistry(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config *bootstrap.Config,
) bootstrap.RouteRegistry {
	routes := make([]Route, 0)
	searchHandler := httpapi_handlers.NewSearch(logger, rdbConn.ToHandle())
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/search",
			handler: searchHandler.Search,
		},
	)

	blocksHandler := httpapi_handlers.NewBlocks(logger, rdbConn.ToHandle())
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/blocks",
			handler: blocksHandler.List,
		},
		Route{
			Method:  GET,
			path:    "api/v1/blocks/{height-or-hash}",
			handler: blocksHandler.FindBy,
		},
		Route{
			Method:  GET,
			path:    "api/v1/blocks/{height}/transactions",
			handler: blocksHandler.ListTransactionsByHeight,
		},
		Route{
			Method:  GET,
			path:    "api/v1/blocks/{height}/events",
			handler: blocksHandler.ListEventsByHeight,
		},
		Route{
			Method:  GET,
			path:    "api/v1/blocks/{height}/commitments",
			handler: blocksHandler.ListCommitmentsByHeight,
		},
	)

	return &RouteRegistry{routes: routes}
}
```

```go
package routes

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	"github.com/valyala/fasthttp"
)

type RouteRegistry struct {
	routes []Route
}

type Route struct {
	Method  string
	path    string
	handler fasthttp.RequestHandler
}

func (registry *RouteRegistry) Register(server *httpapi.Server, routePrefix string) {
	if routePrefix == "/" {
		routePrefix = ""
	}

	for _, route := range registry.routes {
		registerRoute(server, routePrefix, route)
	}
}

func registerRoute(server *httpapi.Server, routePrefix string, route Route) {
	switch route.Method {
	case GET:
		server.GET(fmt.Sprintf("%s/%s", routePrefix, route.path), route.handler)
	}
}

const (
	GET = "GET"
)
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
