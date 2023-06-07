# syntax=docker/dockerfile:1

FROM golang:1.19.0

# Set destination for COPY
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy