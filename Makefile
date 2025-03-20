.PHONY: generate build run clean dev

generate:
	./scripts/generate-templ.sh

build: generate
	go build -o bin/server ./cmd/server

run: build
	./bin/server

# Default example if none specified
# Usage: make dev CMD_DIR=greeting-example
CMD_DIR ?= server
dev:
	EXAMPLE_DIR=$(CMD_DIR) air

clean:
	rm -rf bin/
	find ./templates -name "*_templ.go" -type f -delete
