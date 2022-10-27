.PHONY: has_docker has_docker_compose lint build_ginkgo_image test test_watch ginkgo

GOLANGCI_LINT_VERSION := "v1.49.0"

# Use lazy assignment(`=`) such that command existence are evaluated when used
GO = $(shell command -v go)
GINKGO = $(shell command -v ginkgo)
DOCKER = $(shell command -v docker)
DOCKER_COMPOSE = $(shell command -v docker-compose)

TEST_DB ?= 1

has_golang:
ifndef GO
	@echo "Golang not found. Please install go and try again."
	@exit 1
endif

has_ginkgo:
ifndef GINKGO
	@echo "ginkgo not found. Please install ginkgo with `make install_ginkgo_local` and try again."
	@exit 1
endif

has_docker:
ifndef DOCKER
	@echo "Docker not found. Please install docker and try again."
	@exit 1
endif

has_docker_compose:
ifndef DOCKER_COMPOSE
	@echo "docker-compose not found. Please install docker-compose and try again."
	@exit 1
endif

lint: has_docker has_docker_compose
	docker run --rm -v $(shell pwd):/app -w /app \
		golangci/golangci-lint:${GOLANGCI_LINT_VERSION} \
		golangci-lint run -v --config .golangci.yml

install_ginkgo_local:
ifndef GINKGO
	$(GO) install github.com/onsi/ginkgo/ginkgo@v1.16.5
endif

build_ginkgo_image:
	docker build -t crypto-com/chain-indexing/ginkgo docker/ginkgo

test: has_docker has_docker_compose has_golang build_ginkgo_image
ifeq ($(TEST_DB), 1)
	docker-compose -f docker/docker-compose.test.yml up --abort-on-container-exit
else
	# TODO: Migrate coin test cases to Ginkgo
	docker run --rm -v $(shell pwd):/app -w /app \
		-e TEST_DB=$(TEST_DB) \
		crypto-com/chain-indexing/ginkgo \
		ginkgo -r && \
		go test ./usecase/coin/...
endif

test_watch: has_docker has_docker_compose build_ginkgo_image
	docker run --rm -v $(shell pwd):/app -w /app \
		crypto-com/chain-indexing/ginkgo \
		ginkgo watch -r

ginkgo: has_docker has_docker_compose build_ginkgo_image
	docker run --rm -v $(shell pwd):/app -w /app \
		crypto-com/chain-indexing/ginkgo \
		ginkgo $(filter-out $@,$(MAKECMDGOALS))

test_local: has_golang install_ginkgo_local has_ginkgo
	TEST_DB=$(TEST_DB) $(GINKGO) -r
	$(GO) test ./usecase/coin/...

test_watch_local: has_golang install_ginkgo_local has_ginkgo
	$(GINKGO) watch -r
