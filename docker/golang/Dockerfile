FROM golang:1.13.6-alpine

RUN apk --update add --no-cache git

WORKDIR /go/src/api
COPY ./api /go/src/api

RUN go get -u github.com/lib/pq && \
    go get -u github.com/labstack/echo/... && \
    go get gopkg.in/urfave/cli.v2@master && \
    go get github.com/oxequa/realize
