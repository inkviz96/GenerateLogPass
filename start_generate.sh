#!/bin/bash
if [[ $1 = "-h" ]] || [[ $1 = "--help" ]]
then
    echo "first argument must be a file name like logging.txt"
    echo "second argument must be integer value passwords generation"
    exit
fi

if [ $1 = "[[:alnum:]].txt" ]
then
    echo "File name can't be empty!"
    exit
fi
if [ $2 -lt 1 ]
then
    echo "PassCount must be greate then 0"
    exit
fi
./generate_log_pass -file $1 -passCount $2
