FROM golang:1.20-buster

WORKDIR /userservice/server

COPY controllers/ controllers/
COPY db/ db/
COPY grpc/ grpc/
COPY message_broker/ message_broker/
COPY middlewares/ middlewares/
COPY protos/ protos/
COPY schemas/ schemas/
COPY main.go main.go

COPY go.mod go.mod
RUN go mod tidy
RUN go build || exit 1

ENTRYPOINT [ "./userservice" ]