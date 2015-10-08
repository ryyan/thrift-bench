# thrift-benchmark

Project used to benchmark performance for Thrift in different languages.

## Setup

### Build and run docker image

This could take a few minutes

```
sudo docker build -t thrift-benchmark --no-cache .

sudo docker run -it --rm \
  -v "$PWD"/go:/go/src/go \
  -v "$PWD"/py:/py/py \
  -v "$PWD"/echo.thrift:/echo.thrift \
  thrift-benchmark bash
```

### Build within docker image

```
/go/bin/generator /echo.thrift /go/src/
cd /go/src/go
go get && go build

cd /py/py
python3 -m compileall .
```

## Run benchmarks within docker image

View the current results [here](RESULTS.md)

### Go server & client

```
cd /go/src/go
./go -server &
./go -num 1000
```

### Python server & client

```
cd /py/py
python3 server.py &
python3 client.py 1000
```
