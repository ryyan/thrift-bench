# thrift-benchmark Results

Both Go and Python clients running 10 concurrent clients, and each client makes 100000 requests, for a total of 1,000,000 requests.


## Go server & Go client

```
cd /go/src/go
./go -server > 1.log &
time ./go -num 100000 > 2.log

real	1m16.484s
user	0m0.004s
sys	0m0.000s
```

### Py server & Py client

```
cd /py/py/
python server.pyc > 1.log &
time python client.pyc 100000 > 2.log

real	6m56.417s
user	0m0.000s
sys	0m0.000s
```

### Go server & Py client

```
cd /go/src/go
./go -server > 1.log &
cd /py/py/
rm 2.log
time python client.pyc 100000 > 2.log

real	7m40.482s
user	0m0.000s
sys	0m0.000s
```

### Py server & Go client

```
cd /py/py
python server.pyc > 1.log &
cd /go/src/go
time ./go -num 100000 > 2.log

real	5m18.685s
user	0m0.000s
sys	0m0.000s
```

### Host Info

```
$ docker info
Containers: 26
Images: 18
Storage Driver: aufs
 Root Dir: /var/lib/docker/aufs
 Backing Filesystem: extfs
 Dirs: 72
 Dirperm1 Supported: true
Execution Driver: native-0.2
Logging Driver: json-file
Kernel Version: 3.16.0-4-amd64
Operating System: Debian GNU/Linux 8 (jessie)
CPUs: 2
Total Memory: 1.963 GiB
WARNING: No memory limit support
WARNING: No swap limit support

$ docker version
Client:
 Version:      1.8.2
 API version:  1.20
 Go version:   go1.4.2
 Git commit:   0a8c2e3
 Built:        Thu Sep 10 19:08:05 UTC 2015
 OS/Arch:      linux/amd64

Server:
 Version:      1.8.2
 API version:  1.20
 Go version:   go1.4.2
 Git commit:   0a8c2e3
 Built:        Thu Sep 10 19:08:05 UTC 2015
 OS/Arch:      linux/amd64
```
