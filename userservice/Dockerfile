FROM golang:1.20-buster

WORKDIR /userservice/server
RUN mkdir controllers db middlewares schemas

COPY controllers/ controllers/
COPY db/ db/
COPY grpc/ grpc/
COPY middlewares/ middlewares/
COPY protos/ protos/
COPY schemas/ schemas/
COPY main.go main.go

COPY go.mod go.mod
RUN go mod tidy
RUN go build || exit 1

ENTRYPOINT [ "./userservice" ]