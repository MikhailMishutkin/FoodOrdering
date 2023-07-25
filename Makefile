DB_URL=postgresql://root:root@localhost:5444/restaurant?sslmode=disable
DB_URL=postgresql://root:root@localhost:5445/customer?sslmode=disable
DB_URL=postgresql://root:root@localhost:5446/statistics?sslmode=disable

network:
	docker network create foodordering-network

postgres:
	docker run --name postgres --network foodordering-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root restaurant

dropdb:
	docker exec -it postgres dropdb restaurant

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

http:
	go build -v ./cmd/httpserver

grpc:
	go build -v ./cmd/grpcserver

nats:
	nats-server

.PHONY: network postgres createdb dropdb migrateup migratedown http grpc nats