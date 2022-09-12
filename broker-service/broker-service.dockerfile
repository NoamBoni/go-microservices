FROM golang:1.19-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN chmod +x ./brokerApp

#################################################
FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]