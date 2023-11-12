# syntax=docker/dockerfile:1

## Build
FROM golang:latest AS build

WORKDIR /app

COPY . /app

RUN go mod download

RUN dpkg --add-architecture amd64 \
    && apt update \
    && apt-get install -y --no-install-recommends gcc-x86-64-linux-gnu libc6-dev-amd64-cross

RUN CC=x86_64-linux-gnu-gcc CGO_LDFLAGS=-lrt GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -o /app/main /app/main.go

## Deploy
FROM --platform=linux/amd64 ubuntu

WORKDIR /app

RUN mkdir /app/log

COPY --from=build /app/main /app

ENTRYPOINT ["./main"]