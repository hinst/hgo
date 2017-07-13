#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $DIR
target="fog_app_viewer"
path=$(realpath $target)
echo "$target control: $1"

if [ "$1" = "start" ]
then
	pid=$(pidof $path)
	if [[ $? == 0 ]]
	then
		echo probably already running $path
	else
		nohup $path 2>>error-log.txt 1>>output-log.txt &
	fi
fi

if [ "$1" == "stop" ]
then
	pid=$(pidof $path)
	if [[ $? == 0 ]]
	then
		kill -s SIGINT $pid
		while kill -s 0 $pid
		do 
			sleep 1
		done
		echo stopped
	else
		echo probably not running $path
	fi
fi