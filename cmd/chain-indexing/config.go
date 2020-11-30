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
}

// FileConfig is the struct matches config.toml
type FileConfig struct {
	Blockchain BlockchainConfig
	Sync       SyncConfig
	Tendermint TendermintConfig
	HTTP       HTTPConfig
	Database   DatabaseConfig
	Postgres   PostgresConfig
	Logger     LoggerConfig
}

type BlockchainConfig struct {
	CRODenom  string `toml:"cro_denom"`
	BaseDenom string `toml:"base_denom"`

	AccountAddressPrefix   string `toml:"account_address_prefix"`
	AccountPubKeyPrefix    string `toml:"account_pubkey_prefix"`
	ValidatorAddressPrefix string `toml:"validator_address_prefix"`
	ValidatorPubKeyPrefix  string `toml:"validator_pubkey_prefix"`
	ConNodeAddressPrefix   string `toml:"connode_address_prefix"`
	ConNodePubKeyPrefix    string `toml:"connode_pubkey_prefix"`
}

type SyncConfig struct {
	WindowSize int `toml:"window_size"`
}

type HTTPConfig struct {
	ListeningAddress string `toml:"listening_address"`
}

type TendermintConfig struct {
	HTTPRPCURL string `toml:"http_rpc_url"`
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
