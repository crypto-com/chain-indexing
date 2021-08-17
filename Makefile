all: build install-migrate install
build: chain-indexing check-password-strength
chain-indexing:
	go build ./cmd/chain-indexing/
check-password-strength:
	go build ./cmd/check-password-strength/
install:
	go install ./cmd/chain-indexing/
install-migrate:
	./pgmigrate.sh --install-dependency -- version
migrate:
	./pgmigrate.sh -- -verbose up

.PHONY: docker
docker:
	docker build -t chain-indexing -f Dockerfile .
