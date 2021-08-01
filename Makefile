all: build install
build: chain-indexing check-password-strength
chain-indexing:
	go build ./cmd/chain-indexing/
check-password-strength:
	go build ./cmd/check-password-strength/
install:
	go install ./cmd/chain-indexing/
migrate:
	./pgmigrate.sh -- -verbose up

.PHONY: docker
docker:
	GOOS=linux GOARCH=amd64 go build -o bin/chain-indexing.linux ./cmd/chain-indexing/
	GOOS=linux GOARCH=amd64 go build -o bin/migrate.linux -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate
	docker build -t chain-indexing -f Dockerfile .
