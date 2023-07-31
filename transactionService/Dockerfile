FROM golang:alpine AS builder
LABEL authors="johngus"
WORKDIR transactionsService
COPY ./ ./
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -o transactionsService ./cmd/main.go

CMD /wait && ./transactionsService
