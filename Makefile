MIGRATION_PATH=${shell pwd}/configs/migration
CERT_DIR=./tls
DN_NAME=mock_db
DB_TLS_CA_PATH=$(shell pwd)/tls/mysql/ca-cert.pem
DB_TLS_CLIENT_CERT_PATH=$(shell pwd)/tls/mysql/client-cert.pem
DB_TLS_CLIENT_KEY_PATH=$(shell pwd)/tls/mysql/client-key.pem
DB_USER=mock_user
DB_PASSWORD=mock_password
DB_HOST=0.0.0.0
DB_PORT=3306
# for dev environment
DB_URL=mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DN_NAME}?multiStatements=true

.PHONY: create-migration-file
create-migration-file:
	migrate create -ext sql -dir ${MIGRATION_PATH}/ -seq ${MIGRATE_NAME}

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: mock
mock:
# mock consumer transactions repository
	mockgen -package mock -destination internal/controllers/consumers/repository/mock/mock.go -source=internal/controllers/consumers/repository/repository.go
	mockgen -package mock -destination internal/controllers/consumer_transactions/repository/mock/mock.go -source=internal/controllers/consumer_transactions/repository/repository.go

.PHONY: swagger
swagger:
	swag init -g ./cmd/main.go --output ./docs
	
.PHONY: migrateup
migrateup:
	migrate -path ${MIGRATION_PATH} -database "$(DB_URL)" -verbose up

.PHONY: migratedown
migratedown:
	migrate -path ${MIGRATION_PATH} -database "$(DB_URL)" -verbose down -all

.PHONY: server
server:
	go run ./cmd/main.go --config-file=. --env=production

.PHONY: dev-server-clean
dev-server-clean:
	docker compose -f dev.docker-compose.yml down

.PHONY: dev-server
dev-server:
	docker compose -f dev.docker-compose.yml up -d
	sleep 3
	go run ./cmd/main.go --config-file=./config-dev.yaml --env=develop

# just incase you got error like this: tls: failed to verify certificate: x509: certificate is valid for *.myrepublic.co.id, not storage.googleapis.com
# run below make file command GOPROXY=direct go mod tidy or GOPROXY=direct go get your-go-dependency-url
.PHONY: go-direct
go-direct:
	GOPROXY=direct go mod tidy

.PHONY: certs-clean
certs-clean:
	@find tls/ -type f -not \( -name 'client.cnf' -o -name 'server.cnf' -o -name 'ca.cnf' \) -exec rm -f {} +
	@echo "tls/ path successfully cleaned..."

.PHONY: redis-ca

redis-ca:
	@openssl ecparam -name secp384r1 -genkey -noout -out tls/redis/ca/ca-key.pem > /dev/null 2>&1
	@openssl req -x509 -new -key tls/redis/ca/ca-key.pem -out tls/redis/ca/ca-cert.pem -days 1100 -config tls/redis/ca/ca.cnf > /dev/null 2>&1
	@cat tls/redis/ca/ca-cert.pem tls/redis/ca/ca-key.pem > tls/redis/ca/ca.pem
	@echo "Successfully creating ca cert..."

.PHONY: redis-server-cert
redis-server-cert:
	@openssl ecparam -name secp384r1 -genkey -noout -out tls/redis/server/server-key.pem > /dev/null 2>&1
	@openssl req -new -key tls/redis/server/server-key.pem -out tls/redis/server/server-csr.pem -config tls/redis/server/server.cnf > /dev/null 2>&1
	@openssl x509 -req -in tls/redis/server/server-csr.pem -out tls/redis/server/server-cert.pem -CA tls/redis/ca/ca-cert.pem -CAkey tls/redis/ca/ca-key.pem -CAcreateserial -days 1100 -sha384 -extensions v3_req -extfile tls/redis/server/server.cnf > /dev/null 2>&1
	@echo "Successfully creating server certs..."

.PHONY: redis-client-cert
redis-client-cert:
	@openssl ecparam -name secp384r1 -genkey -noout -out tls/redis/client/client-key.pem > /dev/null 2>&1
	@openssl req -new -key tls/redis/client/client-key.pem -out tls/redis/client/client-csr.pem -config tls/redis/client/client.cnf > /dev/null 2>&1
	@openssl x509 -req -in tls/redis/client/client-csr.pem -out tls/redis/client/client-cert.pem -CA tls/redis/ca/ca-cert.pem -CAkey tls/redis/ca/ca-key.pem -CAcreateserial -days 1100 -sha384 -extensions v3_req -extfile tls/redis/client/client.cnf > /dev/null 2>&1
	@echo "Successfully creating client certs..."

.PHONY: redis-clean
redis-clean:
	@find tls/redis/ -type f -not \( -name 'client.cnf' -o -name 'server.cnf' -o -name 'ca.cnf' \) -exec rm -f {} +
	@echo "tls/redis/ path successfully cleaned..."

.PHONY: redis-serve-cert
redis-serve-cert:
	@mv tls/redis/ca/ca-cert.pem tls/redis/ca-cert.pem
	@mv tls/redis/client/client-cert.pem tls/redis/client-cert.pem
	@mv tls/redis/client/client-key.pem tls/redis/client-key.pem
	@mv tls/redis/server/server-cert.pem tls/redis/server-cert.pem
	@mv tls/redis/server/server-key.pem tls/redis/server-key.pem

.PHONY: redis-cert
redis-cert: redis-clean redis-ca redis-server-cert redis-client-cert redis-serve-cert

.PHONY: mysql-ca
mysql-ca:
	@openssl ecparam -name secp384r1 -genkey -noout -out tls/mysql/ca/ca-key.pem > /dev/null 2>&1
	@openssl req -x509 -new -key tls/mysql/ca/ca-key.pem -out tls/mysql/ca/ca-cert.pem -days 1100 -config tls/mysql/ca/ca.cnf > /dev/null 2>&1
	@cat tls/mysql/ca/ca-cert.pem tls/mysql/ca/ca-key.pem > tls/mysql/ca/ca.pem
	@echo "Successfully creating ca cert..."

.PHONY: mysql-server-cert
mysql-server-cert:
	@openssl ecparam -name secp384r1 -genkey -noout -out tls/mysql/server/server-key.pem > /dev/null 2>&1
	@openssl req -new -key tls/mysql/server/server-key.pem -out tls/mysql/server/server-csr.pem -config tls/mysql/server/server.cnf > /dev/null 2>&1
	@openssl x509 -req -in tls/mysql/server/server-csr.pem -out tls/mysql/server/server-cert.pem -CA tls/mysql/ca/ca-cert.pem -CAkey tls/mysql/ca/ca-key.pem -CAcreateserial -days 1100 -sha384 -extensions v3_req -extfile tls/mysql/server/server.cnf > /dev/null 2>&1
	@echo "Successfully creating server certs..."

.PHONY: mysql-client-cert
mysql-client-cert:
	@openssl ecparam -name secp384r1 -genkey -noout -out tls/mysql/client/client-key.pem > /dev/null 2>&1
	@openssl req -new -key tls/mysql/client/client-key.pem -out tls/mysql/client/client-csr.pem -config tls/mysql/client/client.cnf > /dev/null 2>&1
	@openssl x509 -req -in tls/mysql/client/client-csr.pem -out tls/mysql/client/client-cert.pem -CA tls/mysql/ca/ca-cert.pem -CAkey tls/mysql/ca/ca-key.pem -CAcreateserial -days 1100 -sha384 -extensions v3_req -extfile tls/mysql/client/client.cnf > /dev/null 2>&1
	@echo "Successfully creating client certs..."

.PHONY: mysql-clean
mysql-clean:
	@find tls/mysql/ -type f -not \( -name 'client.cnf' -o -name 'server.cnf' -o -name 'ca.cnf' \) -exec rm -f {} +
	@echo "tls/mysql/ path successfully cleaned..."

.PHONY: mysql-serve-cert
mysql-serve-cert:
	@mv tls/mysql/ca/ca-cert.pem tls/mysql/ca-cert.pem
	@mv tls/mysql/client/client-cert.pem tls/mysql/client-cert.pem
	@mv tls/mysql/client/client-key.pem tls/mysql/client-key.pem
	@mv tls/mysql/server/server-cert.pem tls/mysql/server-cert.pem
	@mv tls/mysql/server/server-key.pem tls/mysql/server-key.pem

.PHONY: mysql-cert
mysql-cert: mysql-clean mysql-ca mysql-server-cert mysql-client-cert mysql-serve-cert