##################
### BASE IMAGE ###
##################

FROM golang:1.18.2-alpine3.16 as build

ARG EXEC_ENV

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify && apk add bash

COPY . .
RUN ./build.sh bin -l $EXEC_ENV


##########################
### ENV-SPECIFIC IMAGE ###
##########################

FROM gcc:12.1 as executor

WORKDIR /app

COPY --from=build /app/srvexec-c-generic ./srvexec-c-generic

EXPOSE 8080

ENV SRVEXEC_LISTEN=0.0.0.0
ENV SRVEXEC_PORT=8080
ENV SRVEXEC_TIMEOUT="5s"
ENTRYPOINT [ "./srvexec-c-generic" ]