#!/bin/bash
generateLog () {
    echo "admin" > "logings.txt"
}
if [[ $1 = "-h" ]] || [[ $1 = "--help" ]]
then
    echo "first argument must be a file name like logging.txt"
    echo "second argument must be integer value passwords generation"
    exit
fi
re='^[0-9]+$'
if [[ $1 =~ $re ]] && [[ -z "$2" ]]
then
    generateLog
    file="logings.txt"
else
    file=$1
fi
if [[ $2 -lt 1 ]] && [[ $2 =~ $re ]]
then
    echo "PassCount must be greate then 0"
    exit
elif [[ $2 =~ $re ]] && [[ $2 -gt 1 ]]
then
    count=$2
else
    count=$1
fi
echo " $file $count"
./generate_log_pass -file $file -passCount $count


