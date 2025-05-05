FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o shop-backend ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/shop-backend .
EXPOSE ${APP_PORT}
CMD ["./shop-backend"]
    