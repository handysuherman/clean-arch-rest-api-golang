#!/bin/sh
set -e

redis-server --tls-port 6379 --port 0 \
    --tls-auth-clients yes \
    --tls-session-caching no \
    --tls-protocols "TLSv1.2 TLSv1.3" \
    --loglevel notice \
    --slowlog-max-len 128 \
    --latency-monitor-threshold 0 \
    --notify-keyspace-events "" \
    --list-max-ziplist-size -2 \
    --activerehashing yes \
    --appendfsync everysec \
    --requirepass mock_redis_password
