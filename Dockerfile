# https://hub.docker.com/_/alpine/
FROM alpine:edge

# Add edge repos for latest packages
RUN rm /etc/apk/repositories && touch /etc/apk/repositories
RUN echo 'http://dl-4.alpinelinux.org/alpine/edge/main/' >> /etc/apk/repositories
RUN echo 'http://dl-4.alpinelinux.org/alpine/edge/community/' >> /etc/apk/repositories

# Install packages and libraries
RUN apk update
RUN apk add -uUf bash git gcc g++ go \
  py-pip python python-dev python3 python3-dev
RUN rm -rf /var/cache/apk/*
RUN pip3 install --upgrade pip
RUN pip3 install cython
RUN pip3 install thriftpy

# Set environment variables
ENV GOPATH /go
