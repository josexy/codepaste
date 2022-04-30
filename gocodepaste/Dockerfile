FROM golang:1.17-stretch AS builder

WORKDIR /build
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o gocodepaste

FROM debian:stretch-slim
WORKDIR /app
COPY --from=builder /build/gocodepaste /app/gocodepaste
COPY ./wait-for.sh /app/wait-for.sh
COPY ./docs /app/docs
COPY ./conf /app/conf
COPY .env /app/.env
COPY ./limiter/lua /app/limiter/lua

RUN chmod +x /app/wait-for.sh

EXPOSE 10086
CMD ["/app/wait-for.sh", "mysql:3306", "--", "/app/gocodepaste", "-c", "conf/config.yaml"]
