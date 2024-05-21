FROM golang:1.20-buster

WORKDIR /statisticsservice/server

COPY message_broker/ message_broker/
COPY schemas/ schemas/
COPY main.go main.go

COPY go.mod go.mod
RUN go mod tidy
RUN go build || exit 1

ENTRYPOINT [ "./statisticsservice" ]