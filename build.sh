#!/bin/bash

PROGNAME=$(basename $0)
SUBCOMMAND=$1
ENVIRONMENTS=()

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

# GOVERS="1.18.3"

# if [[ $(go env GOVERSION) != "go$GOVERS" ]]; then
#     echo "Bad version of go '$(go env GOVERSION)'"
#     exit 1
# fi

envExists() {
    [ ! -f "./environments/$1.go" ] && logerr "Missing environments/$1.go" && return 1
    lang="$(cut -d'-' -f1 <<<$1)"
    # Le proxy n'as pas de fichier de langage
    [ ! -f "./languages/$lang.go" ] && [ $1 != "proxy" ] && logerr "Missing languages/$lang.go" && return 1
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
    echo "e.g.: $PROGNAME bin -l python-generic -l java-generic"
}

sub_container() {
    for complete_env in ${ENVIRONMENTS[@]}; do

        lang="$(cut -d'-' -f1 <<<$complete_env)"
        env="$(cut -d'-' -f2- <<<$complete_env)"

        [ ! -f "./dockerfiles/$complete_env.dockerfile" ] && logerr "./dockerfiles/$complete_env.dockerfile doesn't exists!" && exit 1

        log "Building binary container for ${complete_env}..."
        docker build \
            -t srvexec:bin-${complete_env} \
            --build-arg EXEC_ENV=${complete_env} \
            . || ( logerr "ERROR" && exit 1)

        # if previous command failed, exit
        [ $? -ne 0 ] && exit 1

        log "Building executor container for ${complete_env}..."
        docker build \
            -t srvexec:${complete_env} \
            -f dockerfiles/$complete_env.dockerfile \
            . || ( logerr "ERROR" && exit 1)
    done
}

sub_bin() {
    log "Downloading module..."
    go mod download && go mod verify
    for complete_env in ${ENVIRONMENTS[@]}; do

        lang="$(cut -d'-' -f1 <<<$complete_env)"

        log "Building binary for environment ${complete_env}..."
        CGO_ENABLED=0 go build -v -o srvexec-$complete_env -tags $(to_underscore $complete_env),lib$lang .
    done
    log "Done"
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
    log "Environments: ${ENVIRONMENTS[@]}"
else
    logerr "No environment!"
    exit 1
fi

$RUN 
if [ $? = 127 ]; then
    logerr "Error: '$SUBCOMMAND' is not a known subcommand." >&2
    echo
    sub_help
    exit 1
fi