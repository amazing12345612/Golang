#!/bin/bash

    lsof -i :12344
    lsof -i :12345
    lsof -i :12346
    lsof -i :12347
    lsof -i :12348
    lsof -i :12349
    lsof -i :12350
    lsof -i :12351

    # ip_cmd='lsof -i :12344'
    # ip_status=`$ip_cmd`
    # echo $ip_status
    # array=(${ip_status// / })
    # echo $array
