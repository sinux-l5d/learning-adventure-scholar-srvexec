#!/usr/bin/bash

MODE=$1
LANGUAGE=$2

usage() {
    echo "Usage: $0 <bin|container|dev> <language>"
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

([ -z $MODE ] || [ -z $LANGUAGE ] ) && usage && exit 1

if [ $MODE = "dev" ]; then
    require air
    air -build.cmd="go build -o ./tmp/main -tags $LANGUAGE ."
elif [ $MODE = "bin" ]; then
    go run -tags $LANGUAGE .
    python3 test.py $LANGUAGE
elif [ $MODE = "container" ]; then
    require docker
    docker run --rm --name srvexec-$LANGUAGE -p 8080:8080 srvexec:$LANGUAGE
else
    echo 'Invalid mode!'
    usage
    exit 1
fi