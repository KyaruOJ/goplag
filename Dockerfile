FROM golang:alpine

RUN apk add --no-cache openjdk8-jre git curl; \
  cd /usr/local/lib; \
  curl -O https://www.antlr.org/download/antlr-4.8-complete.jar; \
  cd /usr/bin; \
  curl -O https://code.core.moe/KOJ/goplag/raw/master/shell/antlr4.sh; \
  mv antlr4.sh antlr4; \
  chmod +x antlr4; \
  cd /go; \
  git clone https://code.core.moe/KOJ/goplag.git; \
  cd goplag; \
  go build main.go; \
  mv main /usr/bin/goplag; \
  mkdir -p /goplag

WORKDIR /goplag
