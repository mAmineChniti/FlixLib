all: build-css run-server

build-run-silent: build-css run-silent

build-css:
	@echo "Building CSS"
	bun run build:css

run-server:
	@echo "Running server"
	go run .

run-silent:
	@echo "Running server"
	nohup go run . &