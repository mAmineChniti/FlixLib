all: build-css run-server

test: build-css run-silent


build-css:
	@echo "Building CSS"
	bun run build:css

run-server:
	@echo "Running server"
	air -c .air.toml

run-silent:
	@echo "Running server"
	air -c .air.toml &