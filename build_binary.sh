#!/usr/bin/env bash

CURRENT_VERSION=`git describe`
CGO_ENABLED=0 go build -asmflags -trimpath \
	-ldflags "-s -w -X main.version=$CURRENT_VERSION" cmd/teleirc.go

