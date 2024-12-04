package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"os"

	redisDb "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/databases/redis"
)

func (a *app) redis(ctx context.Context) error {
	opt := redisDb.Config{
		Host:      a.cfg.Databases.Redis.Servers[0],
		Password:  a.cfg.Databases.Redis.Password,
		DB:        a.cfg.Databases.Redis.DB,
		PoolSize:  a.cfg.Databases.Redis.PoolSize,
		TLsEnable: a.cfg.Databases.Redis.EnableTLS,
	}

	if opt.TLsEnable {
		ca_cert, err := os.ReadFile(a.cfg.TLS.Redis.CaPath)
		if err != nil {
			a.log.Fatalf("redis.ca_cert.os.ReadFile.err: %v", err)
			return err
		}

		cert, err := tls.LoadX509KeyPair(a.cfg.TLS.Redis.CertPath, a.cfg.TLS.Redis.KeyPath)
		if err != nil {
			a.log.Fatalf("redis.cert.tls.LoadX509KeyPair.err: %v", err)
			return err
		}

		cert_pool := x509.NewCertPool()
		cert_pool.AppendCertsFromPEM(ca_cert)

		opt.TLs = &tls.Config{
			RootCAs:      cert_pool,
			Certificates: []tls.Certificate{cert},
		}
	}

	conn, err := redisDb.NewUniversalRedisClient(ctx, &opt)
	if err != nil {
		a.log.Warnf("redis.redisDb.NewUniversalRedisClient.err: %v", err)
		return err
	}

	a.redisConnection = conn

	return nil
}
