FROM golang:1.16 AS mod_cacher

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

FROM golang:1.16

COPY --from=mod_cacher $GOPATH/pkg/mod $GOPATH/pkg/mod
COPY . "${GOPATH}/src/github.com/crypto-com/chain-indexing"
RUN cd "${GOPATH}/src/github.com/crypto-com/chain-indexing" \
    && make chain-indexing \
    && go get -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate

WORKDIR /app

RUN cp "${GOPATH}/src/github.com/crypto-com/chain-indexing/chain-indexing" .
COPY ./migrations ./migrations
COPY ./pgmigrate.sh ./migrate

EXPOSE 8080
