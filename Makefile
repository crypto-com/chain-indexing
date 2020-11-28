all:
	go build ./cmd/chain-indexing/
	go install ./cmd/chain-indexing/
migrate:
	./pgmigrate.sh -- -verbose up	
