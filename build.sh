#!/usr/bin/bash

usage() {
    echo "Usage: build.sh -l <languages>"
    echo "e.g. build.sh -l java,python"
    exit $1
}

while getopts hl: flag 
do
    case "${flag}" in
        l) languages="${OPTARG}";;
        h) usage 0;;
        *) usage 1;;
    esac
done


if [ -z "${languages}" ]; then
    usage 1
fi

IFS=',' read -ra LANGUAGES <<< "${languages}"
for lang in ${LANGUAGES[@]}; do
    echo "Building binary container for ${lang}..."
    docker build -t srvexec:bin-${lang} --build-arg EXEC_LANG=${lang} . || ( echo "ERROR" && exit 1)
    echo "Building executor container for ${lang}..."
    docker build -t srvexec:${lang} -f dockerfiles/$lang.dockerfile . || ( echo "ERROR" && exit 1)
done
