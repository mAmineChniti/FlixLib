all: run-server

format-lint: format lint

run-server:
	@echo "Running server"
	air -c .air.toml

run-silent:
	@echo "Running server in silent mode"
	air -c .air.toml &

generate-templ:
	@echo "Generating templates"
	templ generate

format:
	@echo "Formatting go and templ"
	gofmt -l -w .
	templ fmt .

lint:
	@echo "Linting go"
	golangci-lint run .

clean:
	@echo "Cleaning up"
	rm -rf pages/*.go
	rm -rf components/*.go
