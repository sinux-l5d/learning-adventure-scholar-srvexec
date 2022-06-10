FROM golang:1.18.2-alpine3.16 as build

ARG EXEC_LANG

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
# Make the binary static
ENV CGO_ENABLED=0 
RUN go build -v -o srvexec-$EXEC_LANG -tags $EXEC_LANG .


FROM scratch as runtime

ARG EXEC_LANG

COPY --from=build /app/srvexec-$EXEC_LANG /srvexec-$EXEC_LANG

