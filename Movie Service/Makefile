build:
	docker-compose build VK-Test_Ex

run:
	docker-compose up VK-Test_Ex

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:admin@0.0.0.0:5432/postgres?sslmode=disable' up

swag:
	swag init -g cmd/filmoteca/main.go