#!/usr/bin/sh

MODE=$1
ENVIRONMENT=$2
LANGUAGE="$(cut -d'-' -f1 <<<$ENVIRONMENT)"

usage() {
    echo "Usage: $0 <bin|container|dev> <environment>"
    echo "  bin: BUILD and RUN the binary"
    echo "  container: RUN the container (build with ./build.sh)"
    echo "  dev: RUN the development server (hot reloading with Air, see CONTRIBUTING.md)"
}

# Check if the command exists 
require() {
    command -v $1 >/dev/null 2>&1 || {
        echo "Error: $1 is not installed"
        exit 1
    }
}

([ -z $MODE ] || [ -z $ENVIRONMENT ] ) && usage && exit 1

if [ $MODE = "dev" ]; then
    require air
    air -build.cmd="go build -o ./tmp/main -tags $ENVIRONMENT,lib$LANGUAGE ."
elif [ $MODE = "bin" ]; then
    go run -tags $ENVIRONMENT,lib$LANGUAGE .
elif [ $MODE = "container" ]; then
    require docker
    docker run --rm --name srvexec-$ENVIRONMENT -p 8080:8080 srvexec:$ENVIRONMENT
else
    echo 'Invalid mode!'
    usage
    exit 1
fi