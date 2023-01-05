FROM golang:1.19 as base

FROM base as dev

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz --directory $(go env GOPATH)/bin

WORKDIR /opt/api