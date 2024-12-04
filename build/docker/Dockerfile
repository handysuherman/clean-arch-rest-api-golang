# Builder Stage
FROM golang:1.23.3-alpine3.20 AS builder
WORKDIR /app
COPY . .

RUN go build -o main ./cmd/main.go

# Final Stage
FROM alpine:latest
RUN apk update && apk upgrade --available
WORKDIR /app
COPY --from=builder /app/main .
COPY ./scripts/start.sh .
COPY ./scripts/wait-for.sh .

EXPOSE 6880
EXPOSE 7880
EXPOSE 8880
EXPOSE 50050

CMD ["/app/main", "--config-file=.", "--env=production"]
ENTRYPOINT [ "/app/start.sh" ]