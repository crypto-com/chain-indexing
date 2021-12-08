module github.com/crypto-com/chain-indexing

go 1.16

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/brianvoe/gofakeit/v5 v5.10.1
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/calvinlauyh/cosmosutils v0.1.0
	github.com/cenkalti/backoff/v4 v4.0.2
	github.com/cosmos/cosmos-sdk v0.44.2
	github.com/cosmos/ibc-go v1.2.1
	github.com/crypto-org-chain/chain-main/v3 v3.0.0-croeseid
	github.com/crypto-org-chain/cronos v0.6.0-testnet
	github.com/ethereum/go-ethereum v1.10.3
	github.com/ettle/strcase v0.1.1
	github.com/fasthttp/router v1.3.3
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/google/go-querystring v1.0.0
	github.com/google/uuid v1.3.0
	github.com/jackc/pgconn v1.6.4
	github.com/jackc/pgtype v1.4.2
	github.com/jackc/pgx/v4 v4.8.1
	github.com/json-iterator/go v1.1.11
	github.com/lab259/cors v0.2.0
	github.com/luci/go-render v0.0.0-20160219211803-9a04cc21af0f
	github.com/mitchellh/mapstructure v1.4.2
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.17.0
	github.com/peggyjv/gravity-bridge/module v0.2.17
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.25.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.13
	github.com/tharsis/ethermint v0.7.1
	github.com/valyala/fasthttp v1.17.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace github.com/rs/zerolog => github.com/rs/zerolog v1.23.0
