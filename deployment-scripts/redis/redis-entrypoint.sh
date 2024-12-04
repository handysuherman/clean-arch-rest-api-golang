#!/bin/sh
set -e

redis-server --port 6379 \
    --loglevel notice \
    --slowlog-max-len 128 \
    --latency-monitor-threshold 0 \
    --notify-keyspace-events "" \
    --list-max-ziplist-size -2 \
    --activerehashing yes \
    --appendfsync everysec \
    --requirepass mock_redis_password
