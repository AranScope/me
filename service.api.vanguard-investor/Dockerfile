FROM golang:1.13.1

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -o service.api.vanguard-investor
#RUN go install -v ./...

CMD ["./service.api.vanguard-investor"]