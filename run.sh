#!/bin/sh

MODE=$1
ENVIRONMENT=$2
LANGUAGE="$(cut -d'-' -f1 <<<$ENVIRONMENT)"


MAIN="\e[94m"
ERR="\e[31m"
_TEXT="\e[4m"
ENDCOLOR="\e[0m"

log() {
    echo -e "$_TEXT$MAIN$@$ENDCOLOR"
}

logerr() {
    echo -e "$_TEXT$ERR$@$ENDCOLOR"
}

to_underscore() {
    echo "$1" | tr '-' '_'
}

usage() {
    echo "Usage: $0 <bin|container|dev> <environment>"
    echo "  bin: BUILD and RUN the binary"
    echo "  container: RUN the container (build before with ./build.sh)"
    echo "  dev: RUN the development server (hot reloading with Air, see CONTRIBUTING.md)"
}

# Check if the command exists 
require() {
    command -v $1 >/dev/null 2>&1 || {
        logerr "Error: $1 is not installed"
        exit 1
    }
}

([ -z $MODE ] || [ -z $ENVIRONMENT ] ) && usage && exit 1

if [ $MODE = "dev" ]; then
    require air
    require go
    AIR_BUILD_CMD="go build -o ./tmp/main -tags $(to_underscore $ENVIRONMENT),lib$LANGUAGE ."
    log "Running development server with \"$AIR_BUILD_CMD\""
    air -build.cmd="$AIR_BUILD_CMD"
elif [ $MODE = "bin" ]; then
    require go
    log "Running binary with tags $(to_underscore $ENVIRONMENT) and lib$LANGUAGE"
    go run -tags $(to_underscore $ENVIRONMENT),lib$LANGUAGE .
elif [ $MODE = "container" ]; then
    require docker

    log "Building container with $ENVIRONMENT"
    ./build.sh container -l $ENVIRONMENT

    # If the build failed, exit
    [ $? -ne 0 ] && exit 1

    log "Running container with $ENVIRONMENT"
    docker run --rm --name srvexec-$ENVIRONMENT --memory 50mb -p 3005:8080 srvexec:$ENVIRONMENT
else
    echo 'Invalid mode!'
    usage
    exit 1
fi
