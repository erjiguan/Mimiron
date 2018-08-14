#!/bin/sh
FROM golang:latest
MAINTAINER erjiguan "lvyangbupt@163.com"
WORKDIR $GOPATH/src/github.com/erjiguan/Mimiron
ADD . $GOPATH/src/github.com/erjiguan/Mimiron
EXPOSE 9090
RUN cd $GOPATH/src/github.com/erjiguan/Mimiron/cmd/api && go build -o main
ENTRYPOINT ["./cmd/api/main"]