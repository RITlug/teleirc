CURRENT_VERSION=`git describe`
go build -ldflags "-X main.version=$CURRENT_VERSION" cmd/teleirc.go