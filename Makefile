BINARY := radio

.PHONY: build test vet check clean tidy docker-check help

build:
	go build -o $(BINARY) .

test:
	go test ./...

vet:
	go vet ./...

check: vet test build

clean:
	rm -f $(BINARY)

tidy:
	go mod tidy

docker-check:
	docker build -t radio-streamer-ci .
	docker run --rm radio-streamer-ci

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build  Build the $(BINARY) binary"
	@echo "  test   Run tests"
	@echo "  vet    Run go vet"
	@echo "  check         Run vet, test, and build"
	@echo "  docker-check  Run make check in a Linux container (CI parity)"
	@echo "  clean         Remove built binary"
	@echo "  tidy   Run go mod tidy"

.DEFAULT_GOAL := help
