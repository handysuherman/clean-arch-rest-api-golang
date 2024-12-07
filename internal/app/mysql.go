package app

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysqlMigrateInstance "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	internalMySQL "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/databases/mysql"
)

func (a *app) mysql(ctx context.Context) error {
	opt := &internalMySQL.Config{
		Driver: a.cfg.Databases.MySQL.Driver,
		Source: fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true",
			a.cfg.Databases.MySQL.Username,
			a.cfg.Databases.MySQL.Password,
			a.cfg.Databases.MySQL.Host,
			a.cfg.Databases.MySQL.Port,
			a.cfg.Databases.MySQL.DBName,
		),
	}

	if a.cfg.Databases.MySQL.EnableTLS {
		tls, err := a.loadTLsCerts(a.cfg.TLS.MySQL.CaPath, a.cfg.TLS.MySQL.CertPath, a.cfg.TLS.MySQL.KeyPath)
		if err != nil {
			a.log.Warnf("mysql.a.loadTLsCerts.err: %v", err)
			return err
		}

		opt.Source = fmt.Sprintf("%s&%s", opt.Source, "tls=custom")
		opt.TLStype = "custom"
		opt.TLSEnable = a.cfg.Databases.MySQL.EnableTLS
		opt.TLS = tls
	}

	conn, err := internalMySQL.New(opt)
	if err != nil {
		a.log.Warnf("internalMySQL.mysql.New.err: %v", err)
		return err
	}

	a.mysqlConnection = conn

	return nil
}

func (a *app) runDBMigration() error {
	driver, err := mysqlMigrateInstance.WithInstance(a.mysqlConnection, &mysqlMigrateInstance.Config{})
	if err != nil {
		a.log.Warnf("runDBMigration.mysqlMigrateInstance.WithInstance.err: %v", err)
		return err
	}

	migration, err := migrate.NewWithDatabaseInstance(a.cfg.Databases.MySQL.MigrationURL, a.cfg.Databases.MySQL.Driver, driver)
	if err != nil {
		a.log.Warnf("runDBMigration.migrate.NewWithDatabaseInstance.err: %v", err)
		return err
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		a.log.Warnf("runDBMigration.migration.Up.err: %v", err)
		return err
	}

	return nil
}
