FROM golang:1.20

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o http ./cmd/httpserver/main.go
RUN go build -o grpc ./cmd/grpcserver/main.go
RUN go build -o worker1 ./cmd/workers/subscribers/restaurant/main.go
RUN go build -o worker2 ./cmd/workers/subscribers/statistics/main.go