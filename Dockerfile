FROM golang:1.16

WORKDIR /app

COPY ./bin/migrate.linux /go/bin/migrate
COPY ./bin/chain-indexing.linux ./chain-indexing
COPY ./migrations ./migrations
COPY ./pgmigrate.sh ./migrate

EXPOSE 28857