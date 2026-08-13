BINARY := radio

.PHONY: build test vet check clean tidy help

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

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build  Build the $(BINARY) binary"
	@echo "  test   Run tests"
	@echo "  vet    Run go vet"
	@echo "  check  Run vet, test, and build"
	@echo "  clean  Remove built binary"
	@echo "  tidy   Run go mod tidy"

.DEFAULT_GOAL := help
