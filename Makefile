# Makefile for storasd

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=storasd
MODULE_NAME=github.com/backpacker69/storasd

# Source files
SOURCES=$(shell find . -name '*.go' -not -path "./vendor/*")

all: init deps test build

init:
	@if [ ! -f go.mod ]; then \
		echo "Initializing Go module..."; \
		$(GOCMD) mod init $(MODULE_NAME); \
	fi

build: init
	$(GOBUILD) -o $(BINARY_NAME) .

test: init
	$(GOTEST) ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

deps: init
	$(GOGET) -u github.com/boltdb/bolt
	$(GOGET) -u github.com/gin-gonic/gin
	$(GOGET) -u github.com/stretchr/testify
	$(GOCMD) mod tidy

# Linting
lint:
	golangci-lint run

# Generate mocks for testing
mocks:
	mockgen -source=db/db.go -destination=mocks/mock_db.go -package=mocks

# Cross compilation
build-linux: init
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_linux .

build-windows: init
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME).exe .

build-mac: init
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_mac .

# Docker
docker-build:
	docker build -t $(BINARY_NAME) .

docker-run:
	docker run -p 8080:8080 $(BINARY_NAME)

.PHONY: all init build test clean run deps lint mocks build-linux build-windows build-mac docker-build docker-run
