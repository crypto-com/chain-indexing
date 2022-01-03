package pg

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

type ConnConfig struct {
	Host                  string
	Port                  int32
	MaybeUsername         *string
	MaybePassword         *string
	MaybeXMigrationsTable *string
	Database              string
	SSL                   bool
}

type ConnParams struct {
	SSLMode          string `url:"sslmode,omitempty"`
	XMigrationsTable string `url:"x-migrations-table,omitempty"`
}

func (config *ConnConfig) ToURL() string {
	var authStr string
	if config.MaybeUsername != nil || config.MaybePassword != nil {
		authStr = *config.MaybeUsername + ":" + *config.MaybePassword + "@"
	}

	connBaseStr := fmt.Sprintf(
		"postgres://%s%s:%d/%s",
		authStr, config.Host, config.Port, config.Database,
	)

	params := ConnParams{
		SSLMode:          "",
		XMigrationsTable: "",
	}

	if !config.SSL {
		params.SSLMode = "disable"
	}

	if config.MaybeXMigrationsTable != nil {
		params.XMigrationsTable = *config.MaybeXMigrationsTable
	}

	queryValues, queryErr := query.Values(params)
	if queryErr != nil {
		panic(fmt.Sprintf("error constructing Postgres connection string: %v", queryErr))
	}

	queryStr := queryValues.Encode()
	if queryStr == "" {
		return connBaseStr
	}
	return fmt.Sprintf("%s?%s", connBaseStr, queryStr)
}

func (config *ConnConfig) DeepClone() ConnConfig {
	return ConnConfig{
		Host:                  config.Host,
		Port:                  config.Port,
		MaybeUsername:         deepCloneStringPointer(config.MaybeUsername),
		MaybePassword:         deepCloneStringPointer(config.MaybePassword),
		MaybeXMigrationsTable: deepCloneStringPointer(config.MaybeXMigrationsTable),
		Database:              config.Database,
		SSL:                   config.SSL,
	}
}

func deepCloneStringPointer(strPtr *string) *string {
	if strPtr == nil {
		return nil
	}

	strVal := *strPtr
	return &strVal
}
