# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy \
    && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o tiny-http-probe \
    && chown 1000:1000 /app/tiny-http-probe \
    && chmod 755 /app/tiny-http-probe


# Minimal runtime image (from scratch)
FROM scratch
COPY --from=builder /app/tiny-http-probe /tiny-http-probe
EXPOSE 8080
ENV PORT=8080
USER 1000:1000
ENTRYPOINT ["/tiny-http-probe"]