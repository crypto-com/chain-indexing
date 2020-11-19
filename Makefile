all:
	go build ./cmd/chainindex/
	go install ./cmd/chainindex/
migrate:
	./pgmigrate.sh -- -verbose up	
