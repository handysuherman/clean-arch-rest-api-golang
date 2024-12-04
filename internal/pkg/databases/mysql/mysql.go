package mysql

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	// _ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Driver                string         `json:"driver"`
	Source                string         `json:"source"`
	MaxOpenConnection     *int           `json:"max_open_connection"`
	MaxIdleConnection     *int           `json:"max_idle_connection"`
	MaxConnectionIdleTime *time.Duration `json:"max_connection_idle_time"`
	MaxConnectionLifeTime *time.Duration `json:"max_connection_life_time"`
	TLStype               string         `json:"tlsType"`
	TLSEnable             bool           `json:"tlsEnable"`
	TLS                   *tls.Config    `json:"tls"`
}

const (
	maxConn         = 50
	maxConnIdleTime = 1 * time.Minute
	maxConnLifeTime = 3 * time.Minute
)

func New(config *Config) (*sql.DB, error) {
	// TLS
	if config.TLSEnable {
		mysql.RegisterTLSConfig("custom", config.TLS)
	}

	db, err := sql.Open(config.Driver, config.Source)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	// Common settings for both TLS and non-TLS connections
	db.SetMaxIdleConns(10)

	if config.MaxIdleConnection != nil {
		db.SetMaxIdleConns(*config.MaxIdleConnection)
	}

	db.SetMaxOpenConns(maxConn)

	if config.MaxOpenConnection != nil {
		db.SetMaxOpenConns(*config.MaxOpenConnection)
	}

	db.SetConnMaxIdleTime(maxConnIdleTime)

	if config.MaxConnectionIdleTime != nil {
		db.SetConnMaxIdleTime(*config.MaxConnectionIdleTime)
	}

	db.SetConnMaxLifetime(maxConnLifeTime)

	if config.MaxConnectionLifeTime != nil {
		db.SetConnMaxLifetime(*config.MaxConnectionLifeTime)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
