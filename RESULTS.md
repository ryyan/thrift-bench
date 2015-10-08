# thrift-benchmark Results

Both Go and Python clients running 10 concurrent clients, and each client makes 100000 requests, for a total of 1,000,000 requests.


## Go server & Go client

```
Go server & Go client

real	0m40.140s
user	0m0.000s
sys	0m0.000s


Py server & Py client

real	1m19.654s
user	0m0.000s
sys	0m0.000s


Go server & Py client

real	1m0.976s
user	0m0.000s
sys	0m0.000s


Py server & Go client

real	1m23.829s
user	0m0.000s
sys	0m0.000s
```

## Host Info

```
$ docker info
Containers: 6
Images: 35
Storage Driver: aufs
 Root Dir: /var/lib/docker/aufs
 Backing Filesystem: extfs
 Dirs: 47
 Dirperm1 Supported: true
Execution Driver: native-0.2
Logging Driver: json-file
Kernel Version: 3.19.0-30-generic
Operating System: Ubuntu 15.04
CPUs: 4
Total Memory: 3.641 GiB
WARNING: No swap limit support

$ docker version
Client:
 Version:      1.8.2
 API version:  1.20
 Go version:   go1.4.2
 Git commit:   0a8c2e3
 Built:        Thu Sep 10 19:21:21 UTC 2015
 OS/Arch:      linux/amd64

Server:
 Version:      1.8.2
 API version:  1.20
 Go version:   go1.4.2
 Git commit:   0a8c2e3
 Built:        Thu Sep 10 19:21:21 UTC 2015
 OS/Arch:      linux/amd64
```
