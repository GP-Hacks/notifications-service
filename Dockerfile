FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY notifications/go.mod notifications/go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/notifications/cmd/notifications
RUN go build -o notifications_service

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/notifications/cmd/notifications/notifications_service .
COPY --from=builder /app/notifications/config/config.yaml ./config/config.yaml

EXPOSE 8080

CMD ["./notifications_service"]
