FROM golang:1.17 as builder

WORKDIR /httpserver
COPY . .
RUN GOPROXY=https://goproxy.cn make build

FROM ubuntu
WORKDIR /app
COPY --from=builder /httpserver/bin/ ./bin

EXPOSE 8090

ENTRYPOINT ["./bin/amd64/httpserver"]

