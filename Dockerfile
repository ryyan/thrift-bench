FROM alpine:latest

# Install required system packages
RUN apk update && \
  apk add --update  bash git gcc curl mercurial make binutils bison \
  g++ linux-headers nano go python3 && \
  apk upgrade && \
  rm -rf /var/cache/apk/*

ENV GOPATH /go
