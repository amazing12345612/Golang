#!/bin/bash


LISTEN_ADDRESS=127.0.0.1:12345 STORAGE_ROOT=/Users/cjm/tmp/1 go run $1/dataServer/dataServer.go &
LISTEN_ADDRESS=127.0.0.1:12346 STORAGE_ROOT=/Users/cjm/tmp/2 go run $1/dataServer/dataServer.go &
LISTEN_ADDRESS=127.0.0.1:12347 STORAGE_ROOT=/Users/cjm/tmp/3 go run $1/dataServer/dataServer.go &
# LISTEN_ADDRESS=127.0.0.1:12348 STORAGE_ROOT=/Users/cjm/tmp/4 go run $1/dataServer/dataServer.go &
# LISTEN_ADDRESS=127.0.0.1:12349 STORAGE_ROOT=/Users/cjm/tmp/5 go run $1/dataServer/dataServer.go &
# LISTEN_ADDRESS=127.0.0.1:12344 STORAGE_ROOT=/Users/cjm/tmp/6 go run $1/dataServer/dataServer.go &

LISTEN_ADDRESS=127.0.0.1:12350 go run $1/apiServer/apiServer.go &
# LISTEN_ADDRESS=127.0.0.1:12351 go run $1/apiServer/apiServer.go &
