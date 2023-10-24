FROM golang:1.20-alpine AS base

RUN apk update && apk add --no-cache git

WORKDIR /usr/src/app
COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download && go mod verify



COPY . .

RUN go get github.com/nats-io/nats.go
RUN go mod vendor
FROM base AS build-http
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w' -o http ./cmd/httpserver/main.go
#RUN go build -o /bin/http ./cmd/httpserver/main.go
#WORKDIR /app
#ENTRYPOINT [ "./http" ]
FROM base AS build-grpc
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w' -o grpc ./cmd/grpcserver/main.go










#RUN go build -o /bin/grpc ./cmd/grpcserver/main.go
##WORKDIR /app
#ENTRYPOINT [ "./grpc" ]
#FROM base AS build-restaurant
#RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w' -o restaurant ./cmd/workers/subscribers/restaurant/main.go
###RUN go build -o /bin/restaurant ./cmd/workers/subscribers/restaurant/main.go
###WORKDIR /app
##ENTRYPOINT [ "./restaurant" ]
#FROM base AS build-statistics
#RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w' -o statistics ./cmd/workers/subscribers/statistics/main.go
##RUN go build -o /bin/statistics ./cmd/workers/subscribers/statistics/main.go
##WORKDIR /app
#ENTRYPOINT [ "./statistics" ]
#
#FROM alpine:latest AS http
#COPY --from=build-http /bin/http /app/
#WORKDIR /app
#ENTRYPOINT [ "./http" ]
#
#FROM alpine:latest AS grpc
#COPY --from=build-grpc /bin/grpc /app/
#WORKDIR /app
#ENTRYPOINT [ "./grpc" ]
#
#FROM alpine:latest AS restaurant
#COPY --from=build-restaurant /bin/restaurant /app/
#WORKDIR /app
#ENTRYPOINT [ "./restaurant" ]
#
#FROM alpine:latest AS statistics
#COPY --from=build-statistics /bin/statistics /app/
#WORKDIR /app
#ENTRYPOINT [ "./statistics" ]

