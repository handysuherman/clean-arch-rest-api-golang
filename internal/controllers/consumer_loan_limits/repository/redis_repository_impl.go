package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/serializer"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/redis/go-redis/v9"
)

const (
	redisPrefixKey = "mtfinance:consumer_loan_limits"
)

func (r *RedisRepositoryImpl) Put(ctx context.Context, key string, arg *ConsumerLoanLimit) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RedisRepositoryImpl.Put")
	defer span.Finish()

	payload, err := serializer.Marshal(arg)
	if err != nil {
		r.log.Warnf("put.serializer.marshal.err: %w", err)
		return
	}

	prefixKey := helper.RedisPrefixes(
		key,
		redisPrefixKey,
		r.cfg.Databases.Redis.Prefixes.ConsumerLoanLimit.Prefix,
		r.cfg.Databases.Redis.AppID,
	)

	if err := r.redisClient.Set(ctx, prefixKey, payload, r.cfg.Databases.Redis.Prefixes.ConsumerLoanLimit.ExpirationDuration).Err(); err != nil {
		return
	}

	r.log.Debugf("put.prefix: %s, key: %s", prefixKey, key)
}

func (r *RedisRepositoryImpl) Get(ctx context.Context, key string) (*ConsumerLoanLimit, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RedisRepositoryImpl.Get")
	defer span.Finish()

	prefixKey := helper.RedisPrefixes(
		key,
		redisPrefixKey,
		r.cfg.Databases.Redis.Prefixes.ConsumerLoanLimit.Prefix,
		r.cfg.Databases.Redis.AppID,
	)

	msg, err := r.redisClient.Get(ctx, prefixKey).Bytes()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Warnf("redis.client.get.err: %w", err)
		}
		return nil, fmt.Errorf("unable to get cache: %w", tracing.TraceWithError(span, err))
	}

	var payload ConsumerLoanLimit
	if err := serializer.Unmarshal(msg, &payload); err != nil {
		return nil, fmt.Errorf("serializer.unmarshal.err: %w", tracing.TraceWithError(span, err))
	}

	r.log.Debugf("get.prefix: %s, key: %s", prefixKey, key)

	return &payload, nil
}

func (r *RedisRepositoryImpl) Del(ctx context.Context, key string) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RedisRepositoryImpl.Del")
	defer span.Finish()

	prefixKey := helper.RedisPrefixes(
		key,
		redisPrefixKey,
		r.cfg.Databases.Redis.Prefixes.ConsumerLoanLimit.Prefix,
		r.cfg.Databases.Redis.AppID,
	)

	if err := r.redisClient.Del(ctx, prefixKey).Err(); err != nil {
		r.log.Warnf("delete.cache.del.err: %v", err)
		return
	}
	r.log.Debugf("del-prefix: %s, key: %s", prefixKey, key)
}
