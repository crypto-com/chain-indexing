package config

const SYSTEM_MODE_EVENT_STORE = "EVENT_STORE"
const SYSTEM_MODE_TENDERMINT_DIRECT = "TENDERMINT_DIRECT"

type Config struct {
	IndexService  IndexService
	HTTPService   HTTPService
	TendermintApp TendermintApp
	CosmosApp     CosmosApp
	Debug         Debug
	Postgres      Postgres
	Logger        Logger
	Prometheus    Prometheus
}

type IndexService struct {
	Enable                     bool
	Mode                       string
	WindowSize                 int
	Blockchain                 Blockchain
	Projection                 Projection
	CronJob                    CronJob
	CosmosVersionEnabledHeight CosmosVersionEnabledHeight
	GithubAPI                  GithubAPI
}

type HTTPService struct {
	Enable             bool
	ListeningAddress   string
	RoutePrefix        string
	CorsAllowedOrigins []string
	CorsAllowedMethods []string
	CorsAllowedHeaders []string
}

type Blockchain struct {
	BondingDenom           string
	AccountAddressPrefix   string
	AccountPubKeyPrefix    string
	ValidatorAddressPrefix string
	ValidatorPubKeyPrefix  string
	ConNodeAddressPrefix   string
	ConNodePubKeyPrefix    string
}

type Debug struct {
	PprofEnable           bool
	PprofListeningAddress string
}

type TendermintApp struct {
	HTTPRPCUrl           string
	Insecure             bool
	StrictGenesisParsing bool
}

type CosmosApp struct {
	HTTPRPCUrl string
	Insecure   bool
}

type Postgres struct {
	SSL                 bool
	Host                string
	Port                int32
	Username            string
	Password            string
	Name                string
	Schema              string
	MaxConns            int32
	MinConns            int32
	MaxConnLifeTime     string
	MaxConnIdleTime     string
	HealthCheckInterval string
}

type Logger struct {
	Level string
	Color bool
}

type Projection struct {
	Enables      []string
	ExtraConfigs map[string]interface{}
}

type CronJob struct {
	Enables []string
}

type CosmosVersionEnabledHeight struct {
	V0_42_7 uint64
}

type GithubAPI struct {
	Username         string
	Token            string
	MigrationRepoRef string
}

type Prometheus struct {
	Enable     bool
	ExportPath string
	Port       string
}
