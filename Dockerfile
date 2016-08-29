FROM alpine:edge

ENV SHELL /bin/bash
ENV GOPATH /go

WORKDIR /app

RUN echo \
    # Add edge repos for latest packages
    && rm /etc/apk/repositories \
    && touch /etc/apk/repositories \
    && echo 'http://dl-4.alpinelinux.org/alpine/edge/main/' >> /etc/apk/repositories \
    && echo 'http://dl-4.alpinelinux.org/alpine/edge/testing/' >> /etc/apk/repositories \
    && echo 'http://dl-4.alpinelinux.org/alpine/edge/community/' >> /etc/apk/repositories \
    # General
    && apk update \
    && apk add -f bash \
    # Go
    && apk add -f git gcc g++ go \
    && go get github.com/samuel/go-thrift/thrift \
    && go get github.com/samuel/go-thrift/generator \
    # Python
    && apk add -f py-pip python3 python3-dev \
    && pip3 install --upgrade pip \
    && pip3 install cython thriftpy \
    # Rust
    && apk add -f rust cargo \
    && cargo install thrust \
    && export PATH=$PATH:/root/.cargo/bin \
    # Cleanup
    && rm -rf /var/cache/apk/*

CMD "/bin/bash"
