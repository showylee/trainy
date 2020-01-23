FROM golang:1.13.6-alpine

WORKDIR /go/src/app/api
ADD ./api .

RUN apk add git

RUN go get github.com/lib/pq
