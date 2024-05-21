FROM golang:1.20-buster

WORKDIR /taskservice/server

COPY controllers/ controllers/
COPY db/ db/
COPY protos/ protos/
COPY main.go main.go

COPY go.mod go.mod
RUN go mod tidy
RUN go build || exit 1

ENTRYPOINT [ "./taskservice" ]