FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o teleirc cmd/teleirc.go

# Use a minimal base image to run the binary
FROM alpine:latest

RUN adduser -D teleirc-user
USER teleirc-user

COPY --from=builder /app/teleirc /opt/teleirc/teleirc
COPY --from=builder /app/.env /opt/teleirc/conf

WORKDIR /opt/teleirc
ENTRYPOINT [ "./teleirc" ]
CMD [ "-conf", "/opt/teleirc/conf", "-debug", "true" ]
