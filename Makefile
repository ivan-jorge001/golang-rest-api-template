setup:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init --dir cmd/server/
	go build -o bin/server cmd/server/main.go

build:
	docker compose build --no-cache

depedencies:
	docker compose up db -d
	docker compose up redis -d

local: depedencies
	APP_ENV=development go run ./cmd/server/main.go  

up:
	docker compose up -d

down:
	docker compose down

restart:
	docker compose restart

clean:
	docker stop go-rest-api-template
	docker stop dockerPostgres
	docker rm go-rest-api-template
	docker rm dockerPostgres
	docker rm dockerRedis
	docker image rm aitrainer-api-backend
	rm -rf .dbdata
