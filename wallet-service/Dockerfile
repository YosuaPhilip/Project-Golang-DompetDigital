# syntax=docker/dockerfile:1
FROM golang:1.23-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY . .
RUN go build -o /wallet-service ./cmd/server

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /wallet-service /app/wallet-service
EXPOSE 8080
CMD ["/app/wallet-service"]
