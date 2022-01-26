FROM golang:latest

WORKDIR /APP

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build

CMD ["./diction-car"]

EXPOSE 50051