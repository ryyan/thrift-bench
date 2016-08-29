#!/bin/bash
echo 'Building go code'
/go/bin/generator /app/sh/echo.thrift /app/go
cd /app/go
go clean && go get && go build

echo 'Building py code'
cd /app/py
rm -rf __pycache__
rm -rf *.pyc
python3 -m compileall .
