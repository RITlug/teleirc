DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
podman build ${DIR}/../../ -f ${DIR}/Dockerfile -t teleirc:latest
