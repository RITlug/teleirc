DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
docker build ${DIR}/../../ -f ${DIR}/Dockerfile -t teleirc