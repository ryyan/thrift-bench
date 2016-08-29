# thrift-benchmark

Project used to benchmark performance for Thrift in different languages.

## Run

### Build and run docker image

This could take a few minutes. Total size ~600MB.

```
docker build -t thrift-benchmark --no-cache .
docker run -it -v ${PWD}:/app thrift-benchmark bash
```

### Build and run tests within docker image

```
bash /app/sh/run.sh
```

## Results

View the current results [here](RESULTS.md)
