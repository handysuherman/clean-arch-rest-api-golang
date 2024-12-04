package repository

import (
	"context"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/config"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	Put(ctx context.Context, key string, arg *Consumer)
	Get(ctx context.Context, key string) (*Consumer, error)
	Del(ctx context.Context, key string)
}

type RedisRepositoryImpl struct {
	log         logger.Logger
	cfg         *config.Config
	redisClient redis.UniversalClient
}

var _ RedisRepository = (*RedisRepositoryImpl)(nil)

func NewRedisRepositoryImpl(log logger.Logger, cfg *config.Config, redisClient redis.UniversalClient) *RedisRepositoryImpl {
	return &RedisRepositoryImpl{
		log:         log,
		cfg:         cfg,
		redisClient: redisClient,
	}
}
