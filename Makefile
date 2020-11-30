all: build install
build:
	go build ./cmd/chain-indexing/
install:
	go install ./cmd/chain-indexing/
migrate:
	./pgmigrate.sh -- -verbose up	
