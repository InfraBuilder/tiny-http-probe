# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go mod init tiny-http-probe && \
    go get github.com/gofiber/fiber/v2 && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o tiny-http-probe

# Minimal runtime image (from scratch)
FROM scratch
COPY --from=builder /app/tiny-http-probe /tiny-http-probe
EXPOSE 80
ENTRYPOINT ["/tiny-http-probe"]