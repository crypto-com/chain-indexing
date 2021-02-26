module github.com/crypto-com/chain-indexing

go 1.14

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Masterminds/squirrel v1.5.0
	github.com/brianvoe/gofakeit/v5 v5.10.1
	github.com/btcsuite/btcutil v1.0.2
	github.com/calvinlauyh/cosmosutils v0.0.5
	github.com/cosmos/cosmos-sdk v0.41.0
	github.com/fasthttp/router v1.3.3
	github.com/golang-migrate/migrate/v4 v4.12.2
	github.com/google/go-querystring v1.0.0
	github.com/google/uuid v1.1.2
	github.com/jackc/pgconn v1.6.4
	github.com/jackc/pgtype v1.4.2
	github.com/jackc/pgx/v4 v4.8.1
	github.com/json-iterator/go v1.1.10
	github.com/lab259/cors v0.2.0
	github.com/luci/go-render v0.0.0-20160219211803-9a04cc21af0f
	github.com/nbutton23/zxcvbn-go v0.0.0-20201221231540-e56b841a3c88
	github.com/nxadm/tail v1.4.5 // indirect
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.2
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.20.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.3
	github.com/urfave/cli/v2 v2.2.0
	github.com/valyala/fasthttp v1.17.0
	golang.org/x/sys v0.0.0-20201214210602-f9fddec55a1e // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
