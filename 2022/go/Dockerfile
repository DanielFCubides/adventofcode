FROM golang:alpine3.16 AS builder

ENV GO111MODULE=on

ARG DAY

WORKDIR $GOPATH/src/app

COPY ./$DAY/ .

RUN go build -o app

RUN cp app /go/bin

from alpine:latest

ARG DAY

WORKDIR /go/bin

COPY --from=builder /go/bin/app .

COPY ./$DAY/ .

CMD ["./app"]
