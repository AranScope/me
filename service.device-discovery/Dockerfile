FROM golang:1.13.1

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -o main
RUN apt update
RUN apt install -y nmap

CMD ["./main"]

