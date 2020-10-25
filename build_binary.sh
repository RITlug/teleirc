CURRENT_TAG=`git describe --abbrev=0 | tr -d '[:space:]'`
CURRENT_COMMIT=`git rev-list -1 HEAD`
CURRENT_VERSION="${CURRENT_TAG}(${CURRENT_COMMIT})"
DESCRIBE_VERSION=`git describe`
go build -ldflags "-X main.version=$DESCRIBE_VERSION" cmd/teleirc.go