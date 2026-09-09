.PHONY: dev migrate test lint build

dev:
	docker-compose -f infra/docker/docker-compose.dev.yml up

migrate:
	cd apps/api && go run cmd/server/main.go migrate

test:
	cd apps/api && go test ./...
	cd apps/web && npm test

lint:
	cd apps/api && golangci-lint run
	cd apps/web && npm run lint

build:
	docker-compose -f infra/docker/docker-compose.yml build
