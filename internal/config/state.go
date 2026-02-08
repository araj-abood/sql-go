package config

import "araj.com/ar/internal/database"

type State struct {
	Cfg *Config
	Db  *database.Queries
}
