FROM alpine:edge

# Add edge repos for latest packages
RUN rm /etc/apk/repositories && touch /etc/apk/repositories
RUN echo 'http://dl-4.alpinelinux.org/alpine/edge/main/' >> /etc/apk/repositories
RUN echo 'http://dl-4.alpinelinux.org/alpine/edge/community/' >> /etc/apk/repositories

# Install required system packages
RUN apk add -uUf bash git gcc \
  g++ linux-headers nano go python && \
  rm -rf /var/cache/apk/*

ENV GOPATH /go
