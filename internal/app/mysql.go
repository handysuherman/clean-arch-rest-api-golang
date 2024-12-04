package app

import (
	"context"
	"fmt"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/databases/mysql"
)

func (a *app) mysql(ctx context.Context) error {
	opt := &mysql.Config{
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

	conn, err := mysql.New(opt)
	if err != nil {
		a.log.Warnf("mysql.mysql.New.err: %v", err)
		return err
	}

	a.mysqlConnection = conn

	return nil
}
