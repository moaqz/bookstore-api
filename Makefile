.PHONY: docker-up docker-down run docs

# ==============================================================================
# Docker

docker-up:
	@echo "Starting docker environment"
	docker compose up -d

docker-down:
	@echo "Starting docker environment"
	docker compose down

# ==============================================================================
# Main

run:
	@echo 'Running server'
	go run cmd/main.go

# ==============================================================================
# Tools commands

docs:
	@echo 'Generating API documentation using swaggo'
	swag init -g cmd/main.go