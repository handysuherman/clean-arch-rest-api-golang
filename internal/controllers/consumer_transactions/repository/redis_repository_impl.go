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
	redisPrefixKey                  = "mtfinance:consumer_transactions"
	redisPrefixIdempotencyCreateKey = "mtfinance:create_consumer_transactions_idempotency"
	redisPrefixIdempotencyUpdateKey = "mtfinance:update_consumer_transactions_idempotency"
)

func (r *RedisRepositoryImpl) Put(ctx context.Context, key string, arg *ConsumerTransaction) {
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
		r.cfg.Databases.Redis.Prefixes.ConsumerTransaction.Prefix,
		r.cfg.Databases.Redis.AppID,
	)

	if err := r.redisClient.Set(ctx, prefixKey, payload, r.cfg.Databases.Redis.Prefixes.ConsumerTransaction.ExpirationDuration).Err(); err != nil {
		return
	}

	r.log.Debugf("put.prefix: %s, key: %s", prefixKey, key)
}

func (r *RedisRepositoryImpl) Get(ctx context.Context, key string) (*ConsumerTransaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RedisRepositoryImpl.Get")
	defer span.Finish()

	prefixKey := helper.RedisPrefixes(
		key,
		redisPrefixKey,
		r.cfg.Databases.Redis.Prefixes.ConsumerTransaction.Prefix,
		r.cfg.Databases.Redis.AppID,
	)

	msg, err := r.redisClient.Get(ctx, prefixKey).Bytes()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Warnf("redis.client.get.err: %w", err)
		}
		return nil, fmt.Errorf("unable to get cache: %w", tracing.TraceWithError(span, err))
	}

	var payload ConsumerTransaction
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
		r.cfg.Databases.Redis.Prefixes.ConsumerTransaction.Prefix,
		r.cfg.Databases.Redis.AppID,
	)

	if err := r.redisClient.Del(ctx, prefixKey).Err(); err != nil {
		r.log.Warnf("delete.cache.del.err: %v", err)
		return
	}
	r.log.Debugf("del-prefix: %s, key: %s", prefixKey, key)
}

func (r *RedisRepositoryImpl) PutIdempotencyCreate(ctx context.Context, key string, arg int64) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RedisRepositoryImpl.PutIdempotencyCreate")
	defer span.Finish()

	payload, err := serializer.Marshal(arg)
	if err != nil {
		r.log.Warnf("put.serializer.marshal.err: %w", err)
		return
	}

	prefixKey := helper.RedisPrefixes(
		key,
		redisPrefixIdempotencyCreateKey,
		r.cfg.Databases.Redis.Prefixes.CreateConsumerTransactionIdempotency.Prefix,
		r.cfg.Databases.Redis.AppID,
	)

	if err := r.redisClient.Set(ctx, prefixKey, payload, r.cfg.Databases.Redis.Prefixes.CreateConsumerTransactionIdempotency.ExpirationDuration).Err(); err != nil {
		return
	}

	r.log.Debugf("put.prefix: %s, key: %s", prefixKey, key)
}

func (r *RedisRepositoryImpl) GetIdempotencyCreate(ctx context.Context, key string) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RedisRepositoryImpl.GetIdempotencyCreate")
	defer span.Finish()

	prefixKey := helper.RedisPrefixes(
		key,
		redisPrefixIdempotencyCreateKey,
		r.cfg.Databases.Redis.Prefixes.CreateConsumerTransactionIdempotency.Prefix,
		r.cfg.Databases.Redis.AppID,
	)

	msg, err := r.redisClient.Get(ctx, prefixKey).Bytes()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Warnf("redis.client.get.err: %w", err)
		}
		return 0, fmt.Errorf("unable to get cache: %w", tracing.TraceWithError(span, err))
	}

	var payload int64
	if err := serializer.Unmarshal(msg, &payload); err != nil {
		return 0, fmt.Errorf("serializer.unmarshal.err: %w", tracing.TraceWithError(span, err))
	}

	r.log.Debugf("get.prefix: %s, key: %s", prefixKey, key)

	return payload, nil
}

func (r *RedisRepositoryImpl) DelIdempotencyCreate(ctx context.Context, key string) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RedisRepositoryImpl.DelIdempotencyCreate")
	defer span.Finish()

	prefixKey := helper.RedisPrefixes(
		key,
		redisPrefixIdempotencyCreateKey,
		r.cfg.Databases.Redis.Prefixes.CreateConsumerTransactionIdempotency.Prefix,
		r.cfg.Databases.Redis.AppID,
	)

	if err := r.redisClient.Del(ctx, prefixKey).Err(); err != nil {
		r.log.Warnf("delete.cache.del.err: %v", err)
		return
	}
	r.log.Debugf("del-prefix: %s, key: %s", prefixKey, key)
}

func (r *RedisRepositoryImpl) PutIdempotencyUpdate(ctx context.Context, key string, arg int64) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RedisRepositoryImpl.PutIdempotencyUpdate")
	defer span.Finish()

	payload, err := serializer.Marshal(arg)
	if err != nil {
		r.log.Warnf("put.serializer.marshal.err: %w", err)
		return
	}

	prefixKey := helper.RedisPrefixes(
		key,
		redisPrefixIdempotencyUpdateKey,
		r.cfg.Databases.Redis.Prefixes.UpdateConsumerTransactionIdempotency.Prefix,
		r.cfg.Databases.Redis.AppID,
	)

	if err := r.redisClient.Set(ctx, prefixKey, payload, r.cfg.Databases.Redis.Prefixes.UpdateConsumerTransactionIdempotency.ExpirationDuration).Err(); err != nil {
		return
	}

	r.log.Debugf("put.prefix: %s, key: %s", prefixKey, key)
}

func (r *RedisRepositoryImpl) GetIdempotencyUpdate(ctx context.Context, key string) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RedisRepositoryImpl.GetIdempotencyCreate")
	defer span.Finish()

	prefixKey := helper.RedisPrefixes(
		key,
		redisPrefixIdempotencyUpdateKey,
		r.cfg.Databases.Redis.Prefixes.UpdateConsumerTransactionIdempotency.Prefix,
		r.cfg.Databases.Redis.AppID,
	)

	msg, err := r.redisClient.Get(ctx, prefixKey).Bytes()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Warnf("redis.client.get.err: %w", err)
		}
		return 0, fmt.Errorf("unable to get cache: %w", tracing.TraceWithError(span, err))
	}

	var payload int64
	if err := serializer.Unmarshal(msg, &payload); err != nil {
		return 0, fmt.Errorf("serializer.unmarshal.err: %w", tracing.TraceWithError(span, err))
	}

	r.log.Debugf("get.prefix: %s, key: %s", prefixKey, key)

	return payload, nil
}

func (r *RedisRepositoryImpl) DelIdempotencyUpdate(ctx context.Context, key string) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RedisRepositoryImpl.DelIdempotencyCreate")
	defer span.Finish()

	prefixKey := helper.RedisPrefixes(
		key,
		redisPrefixIdempotencyUpdateKey,
		r.cfg.Databases.Redis.Prefixes.UpdateConsumerTransactionIdempotency.Prefix,
		r.cfg.Databases.Redis.AppID,
	)

	if err := r.redisClient.Del(ctx, prefixKey).Err(); err != nil {
		r.log.Warnf("delete.cache.del.err: %v", err)
		return
	}
	r.log.Debugf("del-prefix: %s, key: %s", prefixKey, key)
}
