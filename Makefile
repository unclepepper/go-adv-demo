start:
	go run cmd/main.go

restart: down up

up:
	docker-compose up -d

down:
	docker-compose down  --remove-orphans


m-up:
	go run migrations/auto.go