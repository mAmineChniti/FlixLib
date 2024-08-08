all: generate-templ run-server

generate-templ:
	@echo "Generating templates"
	templ generate

run-server:
	@echo "Running server"
	air -c .air.toml

format:
	@echo "Formatting go and templ"
	go fmt ./...
	templ fmt

clean:
	@echo "Cleaning up"
	rm -rf pages/*.go
	rm -rf components/*.go
