FROM golang:1.18.2-alpine3.16 as build

ARG EXEC_ENV
# ARG EXEC_LANG
# ARG EXEC_NAME=${EXEC_LANG}_${EXEC_ENV}

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
# Make the binary static
# ENV CGO_ENABLED=0 
# RUN go build -v -o srvexec_${EXEC_NAME} -tags ${EXEC_NAME},lib${EXEC_LANG} .

RUN ./build.sh bin -l ${EXEC_ENV}


FROM scratch as runtime

ARG EXEC_ENV
# ARG EXEC_LANG
# ARG EXEC_NAME=${EXEC_LANG}_${EXEC_ENV}

COPY --from=build /app/srvexec-${EXEC_ENV} /srvexec-${EXEC_ENV}

