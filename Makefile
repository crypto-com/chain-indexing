test: go test ./.../

run-ginkgo-docker:
	docker-compose -f docker/docker-compose.test.yml up --abort-on-container-exit
