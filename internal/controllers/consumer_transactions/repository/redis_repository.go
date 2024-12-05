package repository

import (
	"context"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/config"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	PutIdempotencyCreate(ctx context.Context, key string, arg int64)
	GetIdempotencyCreate(ctx context.Context, key string) (int64, error)
	DelIdempotencyCreate(ctx context.Context, key string)

	PutIdempotencyUpdate(ctx context.Context, key string, arg int64)
	GetIdempotencyUpdate(ctx context.Context, key string) (int64, error)
	DelIdempotencyUpdate(ctx context.Context, key string)

	Put(ctx context.Context, key string, arg *ConsumerTransaction)
	Get(ctx context.Context, key string) (*ConsumerTransaction, error)
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
