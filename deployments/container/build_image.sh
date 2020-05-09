DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
go build "${DIR}/../../cmd/teleirc.go"
docker build ${DIR} -t teleirc