#!/bin/sh
set -e

redis-cli --tls \
 --cert /tls/ssl/server-cert.pem \
 --cacert /tls/ssl/ca-cert.pem \
 --key /tls/ssl/server-key.pem \
 -a 'mock_redis_password' \
 -p 6379 \
 --raw incr