#!/usr/bin/sh

PROGNAME=$(basename $0)
SUBCOMMAND=$1
ENVIRONMENTS=()

GOVERS="1.18.3"

if [[ $(go env GOVERSION) != "go$GOVERS" ]]; then
    echo "Bad version of go '$(go env GOVERSION)'"
    exit 1
fi

envExists() {
    [ ! -f "./environments/$1.go" ] && echo "Missing environments/$1.go" && return 1
    lang="$(cut -d'-' -f1 <<<$1)"
    [ ! -f "./languages/$lang.go" ] && echo "Missing languages/$lang.go" && return 1
    return 0
}

to_underscore() {
    echo "$1" | tr '-' '_'
}

to_dash() {
    echo "$1" | tr '_' '-'
}

sub_help() {
    echo "Usage: $PROGNAME (container|bin) -l <environment> [-l <environment2> ...]"
    echo "e.g.: $PROGNAME -l python-generic -l java-generic"
}

sub_container() {
    for complete_env in ${ENVIRONMENTS[@]}; do

        lang="$(cut -d'-' -f1 <<<$complete_env)"
        env="$(cut -d'-' -f2- <<<$complete_env)"

        echo "Building binary container for ${complete_env}..."
        docker build \
            -t srvexec:bin-${complete_env} \
            --build-arg EXEC_ENV=${complete_env} \
            . || ( echo "ERROR" && exit 1)

        [ ! -f "./dockerfiles/$env.dockerfile" ] && echo "./dockerfiles/$complete_env.dockerfile doesn't exists!" && exit 1

        echo "Building executor container for ${complete_env}..."
        docker build \
            -t srvexec:${complete_env} \
            -f dockerfiles/$env.dockerfile \
            . || ( echo "ERROR" && exit 1)
    done
}

sub_bin() {
    echo "Downloading module..."
    go mod download && go mod verify
    for complete_env in ${ENVIRONMENTS[@]}; do

        lang="$(cut -d'-' -f1 <<<$env)"

        echo "Building binary for environment ${complete_env}..."
        CGO_ENABLED=0 go build -v -o srvexec-$complete_env -tags $(to_underscore $complete_env),lib$lang .
    done
    echo "Done"
}


case $SUBCOMMAND in 
    "" | "-h" | "--help" | "help")
        sub_help
        ;;
    *)
        shift
        RUN="sub_${SUBCOMMAND}"
        ;;
esac

while getopts hl: flag 
do
    case "${flag}" in
        l) 
            envExists "$OPTARG" || exit 1
            ENVIRONMENTS+=("$OPTARG");;
        h)
            sub_help
            exit 0
            ;;
        *) 
            sub_help
            exit 1
            ;;
    esac
done

if [ ${#ENVIRONMENTS[@]} -gt 0 ]; then
    echo "Environments: ${ENVIRONMENTS[@]}"
else
    echo "No environment!"
    exit 1
fi

$RUN 
if [ $? = 127 ]; then
    echo "Error: '$SUBCOMMAND' is not a known subcommand." >&2
    echo
    sub_help
    exit 1
fi