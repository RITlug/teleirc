FROM golang:1.15 AS teleirc-builder
WORKDIR /opt/teleirc
COPY go.mod go.sum ./
COPY internal ./internal
COPY cmd ./cmd
RUN go build cmd/teleirc.go

FROM alpine
LABEL maintainer="TeleIRC Team"
COPY --from=teleirc-builder /opt/teleirc/teleirc /opt/teleirc/
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 \
    && addgroup -g 65532 teleirc \
    && adduser -s /bin/bash -h /opt/teleirc -D -H teleirc -u 65532 -G teleirc \
    && chown -R teleirc:teleirc /opt/teleirc
USER teleirc
ENTRYPOINT [ "/opt/teleirc/teleirc" ]