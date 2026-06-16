.PHONY: backend-test backend-run docker-up docker-down

backend-test:
	cd backend && go test ./...

backend-run:
	cd backend && go run ./cmd/api

docker-up:
	cd deployments && docker compose up --build

docker-down:
	cd deployments && docker compose down
