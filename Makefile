all: generate-templ run-server

run-silent: generate-templ run-server-silent

run-server-silent:
	@echo "Running server"
	air -c .air.toml &

generate-templ:
	@echo "Generating templates"
	templ generate

run-server:
	@echo "Running server"
	air -c .air.toml

format:
	@echo "Formatting go and templ"
	gofmt -l -w .
	templ fmt .

clean:
	@echo "Cleaning up"
	rm -rf pages/*.go
	rm -rf components/*.go
