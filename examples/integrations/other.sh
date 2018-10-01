#!/usr/bin/env bash

function test_func(){
	echo "test func $1"
}

echo "wow"  | test_func
