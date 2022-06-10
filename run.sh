#!/usr/bin/bash

LANGUAGE=$1

echo -e "This tool is for testing only\n"

[ -z $LANGUAGE ] && echo "No language!" && exit 1

docker run --rm --name srvexec-$LANGUAGE -p 8080:8080 srvexec:$LANGUAGE