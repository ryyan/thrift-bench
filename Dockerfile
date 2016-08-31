FROM alpine:edge

ENV SHELL /bin/bash
ENV GOPATH /go

WORKDIR /app

RUN echo \
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
    # Cleanup
    && rm -rf /var/cache/apk/*

CMD "/bin/bash"
