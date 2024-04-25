# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOBIN=./bin
BINARY_NAME=univerte

all: build

build:
	$(GOBUILD) -o $(GOBIN)/$(BINARY_NAME) -v ./cmd/univerte

clean:
	$(GOCLEAN)
	rm -rf $(GOBIN)

run: build
	$(GOBIN)/$(BINARY_NAME)

.PHONY: all build test clean run