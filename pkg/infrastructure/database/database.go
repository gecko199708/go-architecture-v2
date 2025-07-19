package database

import (
	"app/pkg/adapter/database"
	"database/sql"
	"errors"

	"github.com/uptrace/bun"
)

type Database struct {
	*bun.DB
}

func Open(config database.Config) (*Database, error) {
	switch config := config.(type) {
	case *database.PostgresConfig:
		sqldb := sql.OpenDB(config.Connector())
		db := bun.NewDB(sqldb, config.Dialect())
		return &Database{db}, nil
	case *database.SQLiteConfig:
		sqldb, err := sql.Open(config.Shim(), config.Conn())
		if err != nil {
			return nil, errors.Join(database.ErrFailureOpenDatabase)
		}
		db := bun.NewDB(sqldb, config.Dialect())
		return &Database{db}, nil
	}
	return nil, database.ErrUnSupportedConfig
}
