FROM golang:1.13.6-alpine

WORKDIR /go/src/app/api
ADD ./api .

RUN apk add --no-cache git && \
    go build -o api && \
    go get github.com/lib/pq && \
    go get -u github.com/labstack/echo/... && \
    go get gopkg.in/urfave/cli.v2@master && \
    go get github.com/oxequa/realize
