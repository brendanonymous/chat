package config

import (
	"chat/pkg/data"
)

// ServerConfig contains parameters, clients, configs, etc. used internally
// This struct is populated in main.go at runtime
type ServerConfig struct {
	DBClient data.DBClientInterface
}
