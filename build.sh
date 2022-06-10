#!/usr/bin/bash

PROGNAME=$(basename $0)
SUBCOMMAND=$1
LANGUAGES=()

GOVERS="1.18.2"

if [[ $(go env GOVERSION) != "go$GOVERS" ]]; then
    echo "Bad version of go '$(go env GOVERSION)'"
    exit 1
fi

sub_help() {
    echo "Usage: $PROGNAME (container|bin) -l <language> [-l <language2...]"
    echo "e.g.: $PROGNAME -l python -l java"
}

sub_container() {
    for lang in ${LANGUAGES[@]}; do
        echo "Building binary container for ${lang}..."
        docker build -t srvexec:bin-${lang} --build-arg EXEC_LANG=${lang} . || ( echo "ERROR" && exit 1)

        [ ! -f "./dockerfiles/$lang.dockerfile" ] && echo "./dockerfiles/$lang.dockerfile doesn't exists!" && exit 1

        echo "Building executor container for ${lang}..."
        docker build -t srvexec:${lang} -f dockerfiles/$lang.dockerfile . || ( echo "ERROR" && exit 1)
    done
}

sub_bin() {
    echo "Downloading module..."
    go mod download && go mod verify
    for lang in ${LANGUAGES[@]}; do
        echo "Building binary for ${lang}..."
        CGO_ENABLED=0 go build -v -o srvexec-$lang -tags $lang .
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
            LANGUAGES+=("$OPTARG");;
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

if [ ${#LANGUAGES[@]} -gt 0 ]; then
    echo "Languages: ${LANGUAGES[@]}"
else
    echo "No language!"
    exit 1
fi

$RUN 
if [ $? = 127 ]; then
    echo "Error: '$SUBCOMMAND' is not a known subcommand." >&2
    echo
    sub_help
    exit 1
fi