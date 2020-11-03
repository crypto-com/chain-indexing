package pg

import "fmt"

type ConnConfig struct {
	Host          string
	Port          uint32
	MaybeUsername *string
	MaybePassword *string
	Database      string
	SSL           bool
}

func (config *ConnConfig) ToURL() string {
	var authStr string
	if config.MaybeUsername != nil || config.MaybePassword != nil {
		authStr = *config.MaybeUsername + ":" + *config.MaybePassword + "@"
	}

	connStr := fmt.Sprintf("postgres://%s%s:%d/%s", authStr, config.Host, config.Port, config.Database)
	if config.SSL {
		return connStr
	} else {
		return connStr + "?sslmode=disable"
	}
}
