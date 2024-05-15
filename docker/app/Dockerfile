FROM golang:1.22.2-bullseye

WORKDIR /app

COPY go.mod go.sum ./

RUN apt-get update \
    && apt-get install git \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/rs/zerolog/log \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install github.com/cosmtrek/air@latest \
    && go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest
