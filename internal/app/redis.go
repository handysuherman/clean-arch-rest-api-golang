package app

import (
	"context"

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
		tls, err := a.loadTLsCerts(a.cfg.TLS.Redis.CaPath, a.cfg.TLS.Redis.CertPath, a.cfg.TLS.Redis.KeyPath)
		if err != nil {
			a.log.Warnf("redis.a.loadTLsCerts.err: %v", err)
			return err
		}

		opt.TLs = tls
	}

	conn, err := redisDb.NewUniversalRedisClient(ctx, &opt)
	if err != nil {
		a.log.Warnf("redis.redisDb.NewUniversalRedisClient.err: %v", err)
		return err
	}

	a.redisConnection = conn

	return nil
}
