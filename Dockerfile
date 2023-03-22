FROM golang:latest

WORKDIR /go/src/app

RUN go mod init teste

COPY . .

RUN go build -o math

CMD ["./math"]