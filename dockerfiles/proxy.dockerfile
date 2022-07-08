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

FROM scratch

WORKDIR /app

COPY --from=build /app/srvexec-proxy ./srvexec-proxy

EXPOSE 8080

ENV SRVEXEC_LISTEN=0.0.0.0
ENV SRVEXEC_PORT=8080

ENTRYPOINT [ "./srvexec-proxy" ]