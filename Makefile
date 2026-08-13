BINARY := radio
IMAGE := radio-streamer

.PHONY: build test vet check clean tidy ci docker-check docker-build help

# Native targets (host Go toolchain; platform-specific audio backend).
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

# CI and cross-platform dev: always run inside Linux container.
ci: docker-check

docker-check:
	docker build -t $(IMAGE) .
	docker run --rm $(IMAGE)

docker-build:
	docker build -t $(IMAGE) .
	docker run --rm -v "$(CURDIR):/out" -w /app $(IMAGE) sh -c "make build && install -m 755 $(BINARY) /out/$(BINARY)"

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  ci            Run vet, test, and build in Docker (same as CI)"
	@echo "  docker-check  Alias for ci"
	@echo "  docker-build  Build the Linux $(BINARY) binary via Docker into the repo root"
	@echo "  build         Build $(BINARY) with the host Go toolchain"
	@echo "  test          Run tests with the host Go toolchain"
	@echo "  vet           Run go vet with the host Go toolchain"
	@echo "  check         Run vet, test, and build on the host"
	@echo "  clean         Remove built binary"
	@echo "  tidy          Run go mod tidy"

.DEFAULT_GOAL := help
