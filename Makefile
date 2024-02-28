all: build-css run-server

test: build-css run-silent

dev: build-css-watch run-server

build-css:
	@echo "Building CSS"
	bun run build:css

build-css-watch:
	@echo "Building CSS"
	bun run build:css-watch

run-server:
	@echo "Running server"
	air -c .air.toml

run-silent:
	@echo "Running server"
	air -c .air.toml &