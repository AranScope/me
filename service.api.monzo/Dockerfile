FROM golang:1.13.1

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -o service.api.monzo

CMD ["./service.api.monzo"]