# Stage 1: Build
FROM golang:1.20-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/app

# Stage 2: Minimal final image
FROM alpine:3.18

# Create non-root user (best practice)
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
COPY --from=builder /app/app /app/

USER appuser

EXPOSE 8080
CMD ["/app/app"]
