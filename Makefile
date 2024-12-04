.PHONY: server
server:
	go run ./cmd/main.go --config-file=. --env=develop

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