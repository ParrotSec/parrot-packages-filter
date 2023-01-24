FROM golang:latest

WORKDIR /go/src/packages-filter

COPY . .

RUN go build -o packages-filter

CMD ["./packages-filter"]