FROM golang:alpine

COPY . /go/src/goplag

RUN cd /go/src/goplag; \
  go build main.go; \
  mv main /usr/bin/goplag; \
  mkdir -p /goplag

WORKDIR /goplag
