# thrift-benchmark Results

Both Go and Python clients running 10 concurrent clients, and each client makes 100000 requests, for a total of 1,000,000 requests.


## Go server & Go client

```
Go server & Go client

real	1m37.771s
user	0m0.000s
sys	0m0.000s


Py server & Py client

real	2m27.169s
user	0m0.000s
sys	0m0.000s


Go server & Py client

real	2m24.881s
user	0m0.000s
sys	0m0.000s


Py server & Go client

real	3m1.109s
user	0m0.000s
sys	0m0.000s
```

## Host Info

```
$ docker info
Containers: 9
Images: 23
Storage Driver: aufs
 Root Dir: /var/lib/docker/aufs
 Backing Filesystem: extfs
 Dirs: 41
 Dirperm1 Supported: true
Execution Driver: native-0.2
Kernel Version: 3.16.0-4-amd64
Operating System: Debian GNU/Linux 8 (jessie)
CPUs: 2
Total Memory: 1.963 GiB
WARNING: No memory limit support
WARNING: No swap limit support

$ docker version
Client version: 1.6.2
Client API version: 1.18
Go version (client): go1.3.3
Git commit (client): 7c8fca2
OS/Arch (client): linux/amd64
Server version: 1.6.2
Server API version: 1.18
Go version (server): go1.3.3
Git commit (server): 7c8fca2
OS/Arch (server): linux/amd64
```
