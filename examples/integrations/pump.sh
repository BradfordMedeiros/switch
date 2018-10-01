#!/usr/bin/env bash

while true; do
	NUM=${RANDOM:0:1}
	if [ "$NUM" -gt "8" ]; then
		echo "pay"
	else 
		echo "push"
	fi
	sleep 0.1
done

