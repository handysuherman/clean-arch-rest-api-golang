This tech stack leverages **Golang Hertz** for API development, **SQLC** for efficient database query generation, **MySQL** for relational data storage, and **Redis** for high-speed caching and key-value storage. Together, they provide a robust and scalable backend system capable of handling complex business logic and high-performance requirements.

Directory Tree
-------------------
      
      build/            contains dockers Dockerfile
      cmd/              project entrypoints
      docs/             contains swagger docs generations { *.yaml, *.json }
      configs/          contains project setups configurations
      tls/              contains tls generation setups
      internal/         project bussiness logic


## Run this project in develop mode using the following command:
~~~
make dev-server
~~~

## Access API Documentation at :
~~~
http://0.0.0.0:50050/docs
~~~
it would automatically redirect you , note that access the server in http://0.0.0.0:50050 for avoiding cors issue.

## Do tests:
~~~
go test ./... -v -count=1 -race -cover
~~~

## Run this project in mTLs mode using the following commands:
### Prepare the self signed certificates:
~~~
make mysql-certs
~~~
~~~
make mysql-serve-cert
~~~
~~~
make redis-cert
~~~
~~~
make redis-serve-cert
~~~

### Set hostname in your hosts, please refer the ./tls/{mysql, redis}/{server, client}*.cnf, for example:
~~~
192.168.x.x constantinopel.xyz-multifinance-redis.server
192.168.x.x constantinopel.xyz-multifinance-mysql.server
~~~

### And Finally Launch the docker-compose for mTLS setup
~~~
make server
~~~

