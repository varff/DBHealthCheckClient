# syntax=docker/dockerfile:1

FROM golang:1.19 AS builder
WORKDIR ./home
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN apt-get update && apt-get install -y apt-utils
RUN apt-get -y install postgresql-client netcat
COPY . ./
RUN go build -o out ./cmd/main.go

CMD [ "./out" ]

