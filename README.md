# thrift-benchmark

Project used to benchmark performance for Thrift in different languages.

## Setup

### Build and run docker image

This could take a few minutes. Total size ~600MB.

```
docker build -t thrift-benchmark --no-cache .

docker run -it --rm \
  -v "$PWD"/go:/go/src/go \
  -v "$PWD"/py:/py/py \
  -v "$PWD"/sh:/sh \
  thrift-benchmark bash
```

### Build and run tests within docker image

```
bash /sh/run
```

View the current results [here](RESULTS.md)

### Bugs

- Mounting volumes might not work when using `docker run` directly on OSX.
