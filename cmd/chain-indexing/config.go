package main

type Config struct {
	FileConfig
}

func (config *Config) OverrideByCLIConfig(cliConfig *CLIConfig) {
	if cliConfig.LogLevel != "" {
		config.Logger.Level = cliConfig.LogLevel
	}
	if cliConfig.LoggerColor != nil {
		config.Logger.Color = *cliConfig.LoggerColor
	}
	if cliConfig.DatabaseSSL != nil {
		config.Database.SSL = *cliConfig.DatabaseSSL
	}
	if cliConfig.DatabaseHost != "" {
		config.Database.Host = cliConfig.DatabaseHost
	}
	if cliConfig.DatabasePort != nil {
		config.Database.Port = *cliConfig.DatabasePort
	}
	if cliConfig.DatabaseUsername != "" {
		config.Database.Username = cliConfig.DatabaseUsername
	}
	// Always overwrite database password with CLI (ENV)
	config.Database.Password = cliConfig.DatabasePassword
	if cliConfig.DatabaseName != "" {
		config.Database.Name = cliConfig.DatabaseName
	}
	if cliConfig.DatabaseSchema != "" {
		config.Database.Schema = cliConfig.DatabaseSchema
	}
	if cliConfig.TendermintHTTPRPCURL != "" {
		config.Tendermint.HTTPRPCURL = cliConfig.TendermintHTTPRPCURL
	}
	if cliConfig.CosmosHTTPRPCURL != "" {
		config.CosmosApp.HTTPRPCUL = cliConfig.CosmosHTTPRPCURL
	}
}

type CLIConfig struct {
	LoggerColor *bool
	LogLevel    string

	DatabaseSSL      *bool
	DatabaseHost     string
	DatabasePort     *int32
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
	DatabaseSchema   string

	TendermintHTTPRPCURL string
	CosmosHTTPRPCURL     string
}

// FileConfig is the struct matches config.toml
type FileConfig struct {
	Blockchain BlockchainConfig
	System     SystemConfig
	Sync       SyncConfig
	Tendermint TendermintConfig
	CosmosApp  CosmosAppConfig `toml:"cosmosapp"`
	HTTP       HTTPConfig
	Debug      DebugConfig
	Database   DatabaseConfig
	Postgres   PostgresConfig
	Logger     LoggerConfig
	Projection ProjectionConfig
	Crossfire  CrossfireConfig
}

type BlockchainConfig struct {
	BondingDenom           string `toml:"bonding_denom"`
	AccountAddressPrefix   string `toml:"account_address_prefix"`
	AccountPubKeyPrefix    string `toml:"account_pubkey_prefix"`
	ValidatorAddressPrefix string `toml:"validator_address_prefix"`
	ValidatorPubKeyPrefix  string `toml:"validator_pubkey_prefix"`
	ConNodeAddressPrefix   string `toml:"connode_address_prefix"`
	ConNodePubKeyPrefix    string `toml:"connode_pubkey_prefix"`
}

type SystemConfig struct {
	Mode string `toml:"mode"`
}

type SyncConfig struct {
	WindowSize int `toml:"window_size"`
}

type HTTPConfig struct {
	ListeningAddress   string   `toml:"listening_address"`
	RoutePrefix        string   `toml:"route_prefix"`
	CorsAllowedOrigins []string `toml:"cors_allowed_origins"`
	CorsAllowedMethods []string `toml:"cors_allowed_methods"`
	CorsAllowedHeaders []string `toml:"cors_allowed_headers"`
}

type DebugConfig struct {
	PprofEnable           bool   `toml:"pprof_enable"`
	PprofListeningAddress string `toml:"pprof_listening_address"`
}

type TendermintConfig struct {
	HTTPRPCURL string `toml:"http_rpc_url"`
}

type CosmosAppConfig struct {
	HTTPRPCUL string `toml:"http_rpc_url"`
}

type DatabaseConfig struct {
	SSL      bool   `toml:"ssl"`
	Host     string `toml:"host"`
	Port     int32  `toml:"port"`
	Username string `toml:"username"`
	Password string
	Name     string `toml:"name"`
	Schema   string `toml:"schema"`
}

type PostgresConfig struct {
	MaxConns            int32  `toml:"pool_max_conns"`
	MinConns            int32  `toml:"pool_min_conns"`
	MaxConnLifeTime     string `toml:"pool_max_conn_lifetime"`
	MaxConnIdleTime     string `toml:"pool_max_conn_idle_time"`
	HealthCheckInterval string `toml:"pool_health_check_interval"`
}

type LoggerConfig struct {
	Level string `toml:"level"`
	Color bool   `toml:"color"`
}

type ProjectionConfig struct {
	Enables []string `toml:"enables"`
}

type CrossfireConfig struct {
	PhaseOneStartTime        int64  `toml:"phase_one_start_time"`
	PhaseTwoStartTime        int64  `toml:"phase_two_start_time"`
	PhaseThreeStartTime      int64  `toml:"phase_three_start_time"`
	CompetitionEndTime       int64  `toml:"competition_end_time"`
	JackpotOneStartTime      int64  `toml:"jackpot_one_start_time"`
	JackpotTwoStartTime      int64  `toml:"jackpot_two_start_time"`
	JackpotThreeStartTime    int64  `toml:"jackpot_three_start_time"`
	JackpotFourStartTime     int64  `toml:"jackpot_four_start_time"`
	JackpotFourEndTime       int64  `toml:"jackpot_four_end_time"`
	AdminAddress             string `toml:"admin_address"`
	NetworkUpgradeProposalID string `toml:"network_upgrade_proposal_id"`
	ParticipantsListURL      string `toml:"participants_list_url"`
}
