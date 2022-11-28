#!/bin/bash

port_be_kill=("12344","12345")
funWithReturn(){
    for i in port_be_kill; 
    do
    ip_cmd='lsof -i tcp:'$i
    ip_status=`$ip_cmd`
    done
}
funWithReturn
    # lsof -i :12344
    # lsof -i :12345
    # lsof -i :12346
    # lsof -i :12347
    # lsof -i :12348
    # lsof -i :12349
    # lsof -i :12350
    # lsof -i :12351