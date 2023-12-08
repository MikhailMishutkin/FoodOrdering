migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down -1

http:
	go build -v ./cmd/httpserver

grpc:
	go build -v ./cmd/grpcserver

subscribers:
	go build -v ./cmd/workers
	#go build -v ./cmd/workers

build: grpc http subscribers

nats:
	nats-server

db:
	docker-compose up -d --build