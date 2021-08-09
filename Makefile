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
	cp .gitignore .dockerignore && echo "**/.git" >> .dockerignore
	docker build -t chain-indexing -f Dockerfile .
