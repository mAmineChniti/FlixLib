all: generate-templ run-server

generate-templ:
	@echo "Generating templates"
	templ generate

run-server:
	@echo "Running server"
	air -c .air.toml
