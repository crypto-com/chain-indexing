package config

const SYSTEM_MODE_EVENT_STORE = "EVENT_STORE"
const SYSTEM_MODE_TENDERMINT_DIRECT = "TENDERMINT_DIRECT"

type Config struct {
	Blockchain    Blockchain    `yaml:"blockchain" toml:"blockchain" xml:"blockchain" json:"blockchain"`
	IndexService  IndexService  `yaml:"index_service" toml:"index_service" xml:"index_service" json:"index_service"`
	HTTPService   HTTPService   `yaml:"http_service" toml:"http_service" xml:"http_service" json:"http_service"`
	TendermintApp TendermintApp `yaml:"tendermint_app" toml:"tendermint_app" xml:"tendermint_app" json:"tendermint_app"`
	CosmosApp     CosmosApp     `yaml:"cosmos_app" toml:"cosmos_app" xml:"cosmos_app" json:"cosmos_app"`
	Debug         Debug         `yaml:"debug" toml:"debug" xml:"debug" json:"debug"`
	Postgres      Postgres      `yaml:"postgres" toml:"postgres" xml:"postgres" json:"postgres"`
	Logger        Logger        `yaml:"logger" toml:"logger" xml:"logger" json:"logger"`
	Prometheus    Prometheus    `yaml:"prometheus" toml:"prometheus" xml:"prometheus" json:"prometheus"`
}

type IndexService struct {
	Enable                     bool                       `yaml:"enable" toml:"enable" xml:"enable" json:"enable,omitempty"`
	Mode                       string                     `yaml:"mode" toml:"mode" xml:"mode" json:"mode,omitempty"`
	WindowSize                 int                        `yaml:"window_size" toml:"window_size" xml:"window_size" json:"window_size,omitempty"`
	Projection                 Projection                 `yaml:"projection" toml:"projection" xml:"projection" json:"projection"`
	CronJob                    CronJob                    `yaml:"cron_job" toml:"cron_job" xml:"cron_job" json:"cron_job"`
	CosmosVersionEnabledHeight CosmosVersionEnabledHeight `yaml:"cosmos_version_enabled_height" toml:"cosmos_version_enabled_height" xml:"cosmos_version_enabled_height" json:"cosmos_version_enabled_height"`
	GithubAPI                  GithubAPI                  `yaml:"github_api" toml:"github_api" xml:"github_api" json:"github_api"`
}

type HTTPService struct {
	Enable             bool     `yaml:"enable" toml:"enable" xml:"enable" json:"enable,omitempty"`
	ListeningAddress   string   `yaml:"listening_address" toml:"listening_address" xml:"listening_address" json:"listening_address,omitempty"`
	RoutePrefix        string   `yaml:"route_prefix" toml:"route_prefix" xml:"route_prefix" json:"route_prefix,omitempty"`
	CorsAllowedOrigins []string `yaml:"cors_allowed_origins" toml:"cors_allowed_origins" xml:"cors_allowed_origins" json:"cors_allowed_origins,omitempty"`
	CorsAllowedMethods []string `yaml:"cors_allowed_methods" toml:"cors_allowed_methods" xml:"cors_allowed_methods" json:"cors_allowed_methods,omitempty"`
	CorsAllowedHeaders []string `yaml:"cors_allowed_headers" toml:"cors_allowed_headers" xml:"cors_allowed_headers" json:"cors_allowed_headers,omitempty"`
}

type Blockchain struct {
	BondingDenom           string `yaml:"bonding_denom" toml:"bonding_denom" xml:"bonding_denom" json:"bonding_denom,omitempty"`
	AccountAddressPrefix   string `yaml:"account_address_prefix" toml:"account_address_prefix" xml:"account_address_prefix" json:"account_address_prefix,omitempty"`
	AccountPubKeyPrefix    string `yaml:"account_pub_key_prefix" toml:"account_pub_key_prefix" xml:"account_pub_key_prefix" json:"account_pub_key_prefix,omitempty"`
	ValidatorAddressPrefix string `yaml:"validator_address_prefix" toml:"validator_address_prefix" xml:"validator_address_prefix" json:"validator_address_prefix,omitempty"`
	ValidatorPubKeyPrefix  string `yaml:"validator_pub_key_prefix" toml:"validator_pub_key_prefix" xml:"validator_pub_key_prefix" json:"validator_pub_key_prefix,omitempty"`
	ConNodeAddressPrefix   string `yaml:"con_node_address_prefix" toml:"con_node_address_prefix" xml:"con_node_address_prefix" json:"con_node_address_prefix,omitempty"`
	ConNodePubKeyPrefix    string `yaml:"con_node_pub_key_prefix" toml:"con_node_pub_key_prefix" xml:"con_node_pub_key_prefix" json:"con_node_pub_key_prefix,omitempty"`
}

type Debug struct {
	PprofEnable           bool   `yaml:"pprof_enable" toml:"pprof_enable" xml:"pprof_enable" json:"pprof_enable,omitempty"`
	PprofListeningAddress string `yaml:"pprof_listening_address" toml:"pprof_listening_address" xml:"pprof_listening_address" json:"pprof_listening_address,omitempty"`
}

type TendermintApp struct {
	HTTPRPCUrl           string `yaml:"http_rpc_url" toml:"http_rpc_url" xml:"http_rpc_url" json:"http_rpc_url,omitempty"`
	Insecure             bool   `yaml:"insecure" toml:"insecure" xml:"insecure" json:"insecure,omitempty"`
	StrictGenesisParsing bool   `yaml:"strict_genesis_parsing" toml:"strict_genesis_parsing" xml:"strict_genesis_parsing" json:"strict_genesis_parsing,omitempty"`
}

type CosmosApp struct {
	HTTPRPCUrl string `yaml:"http_rpc_url" toml:"http_rpc_url" xml:"http_rpc_url" json:"http_rpc_url,omitempty"`
	Insecure   bool   `yaml:"insecure" toml:"insecure" xml:"insecure" json:"insecure,omitempty"`
}

type Postgres struct {
	SSL                     bool   `yaml:"ssl" toml:"ssl" xml:"ssl" json:"ssl,omitempty"`
	Host                    string `yaml:"host" toml:"host" xml:"host" json:"host,omitempty"`
	Port                    int32  `yaml:"port" toml:"port" xml:"port" json:"port,omitempty"`
	Username                string `yaml:"username" toml:"username" xml:"username" json:"username,omitempty"`
	Password                string `yaml:"password" toml:"password" xml:"password" json:"password,omitempty"`
	Name                    string `yaml:"name" toml:"name" xml:"name" json:"name,omitempty"`
	Schema                  string `yaml:"schema" toml:"schema" xml:"schema" json:"schema,omitempty"`
	PoolMaxConns            int32  `yaml:"pool_max_conns" toml:"pool_max_conns" xml:"pool_max_conns" json:"pool_max_conns,omitempty"`
	PoolMinConns            int32  `yaml:"pool_min_conns" toml:"pool_min_conns" xml:"pool_min_conns" json:"pool_min_conns,omitempty"`
	PoolMaxConnLifeTime     string `yaml:"pool_max_conn_life_time" toml:"pool_max_conn_life_time" xml:"pool_max_conn_life_time" json:"pool_max_conn_life_time,omitempty"`
	PoolMaxConnIdleTime     string `yaml:"pool_max_conn_idle_time" toml:"pool_max_conn_idle_time" xml:"pool_max_conn_idle_time" json:"pool_max_conn_idle_time,omitempty"`
	PoolHealthCheckInterval string `yaml:"pool_health_check_interval" toml:"pool_health_check_interval" xml:"pool_health_check_interval" json:"pool_health_check_interval,omitempty"`
}

type Logger struct {
	Level string `yaml:"level" toml:"level" xml:"level" json:"level,omitempty"`
	Color bool   `yaml:"color" toml:"color" xml:"color" json:"color,omitempty"`
}

type Projection struct {
	Enables      []string               `yaml:"enables" toml:"enables" xml:"enables" json:"enables,omitempty"`
	ExtraConfigs map[string]interface{} `yaml:"extra_configs" toml:"extra_configs" xml:"extra_configs" json:"extra_configs,omitempty"`
}

type CronJob struct {
	Enables      []string               `yaml:"enables" toml:"enables" xml:"enables" json:"enables,omitempty"`
	ExtraConfigs map[string]interface{} `yaml:"extra_configs" toml:"extra_configs" xml:"extra_configs" json:"extra_configs,omitempty"`
}

type CosmosVersionEnabledHeight struct {
	V0_42_7 uint64 `yaml:"v_0_42_7" toml:"v_0_42_7" xml:"v_0_42_7" json:"v_0_42_7,omitempty"`
}

type GithubAPI struct {
	Username         string `yaml:"username" toml:"username" xml:"username" json:"username,omitempty"`
	Token            string `yaml:"token" toml:"token" xml:"token" json:"token,omitempty"`
	MigrationRepoRef string `yaml:"migration_repo_ref" toml:"migration_repo_ref" xml:"migration_repo_ref" json:"migration_repo_ref,omitempty"`
}

type Prometheus struct {
	Enable     bool   `yaml:"enable" toml:"enable" xml:"enable" json:"enable,omitempty"`
	ExportPath string `yaml:"export_path" toml:"export_path" xml:"export_path" json:"export_path,omitempty"`
	Port       string `yaml:"port" toml:"port" xml:"port" json:"port,omitempty"`
}
