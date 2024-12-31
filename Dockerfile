# Use the official Golang image to build the application
FROM golang:1.23.4-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./main ./src/common/infrastructure/http/server/main.go

# Start a new stage from scratch
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/environment ./environment
EXPOSE 3000
CMD ["./main"]