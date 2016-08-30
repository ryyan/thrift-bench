#!/bin/bash
echo 'Building go'
$GOPATH/bin/generator /app/sh/echo.thrift $GOPATH/src
cd /app/go
go clean && go get -d && go build

echo 'Building py'
cd /app/py
rm -rf __pycache__ > /dev/null
rm -rf *.pyc > /dev/null
python3 -m compileall . > /dev/null
