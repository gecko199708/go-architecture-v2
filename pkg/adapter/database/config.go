package database

import (
	"database/sql/driver"

	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/schema"
)

type Config interface {
	Driver() string
}

type PostgresConfig struct {
}

func (c *PostgresConfig) Driver() string {
	return "postgres"
}

func (c *PostgresConfig) Connector() driver.Connector {
	// TODO: implement
	return pgdriver.NewConnector(pgdriver.WithDSN(""))
}

func (c *PostgresConfig) Dialect() schema.Dialect {
	return pgdialect.New()
}

type SQLiteConfig struct {
	Path string
}

func (c *SQLiteConfig) Driver() string {
	return "sqlite3"
}

func (c *SQLiteConfig) Conn() string {
	// TODO: implement
	return c.Path
}

func (c *SQLiteConfig) Shim() string {
	return sqliteshim.ShimName
}

func (c *SQLiteConfig) Dialect() schema.Dialect {
	return sqlitedialect.New()
}
