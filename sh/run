#!/bin/bash
bash /sh/setup
sleep 1 && pkill -f server

sleep 2 && echo
echo 'Go server & Go client'
cd /go/src/go
./go -server &
sleep 3
time ./go -num 100000
sleep 1 && pkill -f server

sleep 2 && echo
echo 'Py server & Py client'
cd /py/py/
python3 server.py &
sleep 3
time python3 client.py 100000
sleep 1 && sleep 1 && pkill -f server

sleep 2 && echo
echo 'Go server & Py client'
cd /go/src/go
./go -server &
sleep 3
cd /py/py/
time python3 client.py 100000
sleep 1 && pkill -f server

sleep 2 && echo
echo 'Py server & Go client'
cd /py/py
python3 server.py &
sleep 3
cd /go/src/go
time ./go -num 100000
sleep 1 && pkill -f server
