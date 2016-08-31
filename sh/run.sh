#!/bin/bash

function kill_server {
  # Kill any currently running thrift servers
  pkill -f server && sleep 6
}

function run_go_server {
  echo '# Go server'
  kill_server
  cd /app/go
  ./go -server &
}

function run_py_server {
  echo '# Py server'
  kill_server
  cd /app/py
  python3 server.py &
}

function run_clients {
  echo '## Go client' && sleep 10
  cd /app/go
  time ./go -num 100000

  echo '## Py client' && sleep 10
  cd /app/py
  time python3 client.py 100000
}

# Setup
bash /app/sh/setup.sh

# Go server tests
run_go_server
run_clients

# Python server tests
run_py_server
run_clients
