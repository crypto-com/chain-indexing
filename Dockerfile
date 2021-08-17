FROM golang:1.16 AS mod_cacher

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

FROM golang:1.16 AS builder

COPY --from=mod_cacher $GOPATH/pkg/mod $GOPATH/pkg/mod
COPY . "${GOPATH}/src/github.com/crypto-com/chain-indexing"
RUN cd "${GOPATH}/src/github.com/crypto-com/chain-indexing" \
    && make install-migrate \
    && make chain-indexing

FROM golang:1.16

WORKDIR /app

COPY --from=builder "${GOPATH}/bin/migrate" "${GOPATH}/bin/migrate"
COPY --from=builder "${GOPATH}/src/github.com/crypto-com/chain-indexing/chain-indexing" .
COPY ./migrations ./migrations
COPY ./pgmigrate.sh ./migrate
