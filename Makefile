build:
	docker compose up -d --build

down:
	docker compose down

dev-run:
	go run ./pkg/main/main.go


