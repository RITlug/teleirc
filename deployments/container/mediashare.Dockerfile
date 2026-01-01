FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o mediashare ./cmd/mediashare

# Use a minimal base image to run the binary
FROM alpine:latest

RUN adduser -D mediashare-user && \
    mkdir -p /app/data/uploads && \
    chown -R mediashare-user:mediashare-user /app

USER mediashare-user

COPY --from=builder /app/mediashare /app/mediashare

WORKDIR /app

# Default environment variables
ENV MEDIASHARE_PORT=8090
ENV MEDIASHARE_STORAGE_PATH=/app/data/uploads
ENV MEDIASHARE_DB_PATH=/app/data/mediashare.db
ENV MEDIASHARE_RETENTION_HOURS=168
ENV MEDIASHARE_LANG=en

EXPOSE 8090

VOLUME ["/app/data"]

ENTRYPOINT ["./mediashare"]
