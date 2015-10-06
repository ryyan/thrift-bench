# thrift-benchmark

Project used to gather performance benchmarks for Thrift in different languages.

## Prereqs

- Install [docker](https://docs.docker.com/)
- Install [thrift](https://thrift.apache.org/docs/install/)
  - Optional. Used if you want to re-generate the thrift files (gen-go, gen-py, etc)

## Setup

### Generate thrift files

```
thrift -r --gen go echo.thrift
thrift -r --gen py echo.thrift
```

### Build and run docker image

```
sudo docker build -t thrift-benchmark --no-cache .
sudo docker run -it --rm \
  -v "$PWD"/go:/go/src/go \
  -v "$PWD"/gen-go/echo:/go/src/echo \
  -v "$PWD"/py:/py/py \
  -v "$PWD"/gen-py/echo:/py/echo \
  thrift-benchmark bash
```

### Build within docker image

```
cd /go/src/go
go get
go build

cd /py
wget -O thriftpy https://pypi.python.org/packages/source/t/thrift/thrift-0.9.2.tar.gz#md5=91f1c224c46a257bb428431943387dfd
tar -zxf thriftpy && rm thriftpy
cd thrift-0.9.2 && python setup.py install
cd /py
python -m compileall -f .
```

## Run benchmarks within docker image

### Go server & Go client

```
cd /go/src/go
./go -server &
./go -num 1000
```
