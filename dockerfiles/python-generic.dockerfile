##################
### BASE IMAGE ###
##################

FROM golang:1.18.2-alpine3.16 as build

ARG EXEC_ENV=python-generic

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify && apk add bash

COPY . .
RUN ./build.sh bin -l $EXEC_ENV


##########################
### ENV-SPECIFIC IMAGE ###
##########################

FROM python:3.10-slim as executor

WORKDIR /app

COPY --from=build /app/srvexec-python-generic ./srvexec-python-generic

EXPOSE 8080

ENV SRVEXEC_LISTEN=0.0.0.0
ENV SRVEXEC_PORT=8080
ENV SRVEXEC_TIMEOUT="5s"
ENTRYPOINT [ "./srvexec-python-generic" ]
