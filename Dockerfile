FROM golang:alpine

RUN apk add --no-cache git curl; \
  cd /go/src; \
  git clone https://code.core.moe/KOJ/goplag.git; \
  cd goplag; \
  go build main.go; \
  mv main /usr/bin/goplag; \
  mkdir -p /goplag

WORKDIR /goplag
