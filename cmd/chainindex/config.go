package main

// FileConfig is the struct matches config.toml
type FileConfig struct {
	Blockchain BlockchainConfig
	Sync       SyncConfig
	Tendermint TendermintConfig
	HTTP       HTTPConfig
	Database   DatabaseConfig
	Postgres   PostgresConfig
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
	Port     uint32 `toml:"port"`
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
