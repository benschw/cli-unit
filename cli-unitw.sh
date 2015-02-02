#!/bin/bash


BIN_PATH=".cli-unit/cli-unit"

if [ ! -f $BIN_PATH ]; then
	mkdir -p $(dirname $BIN_PATH)
    curl -sS https://drone.io/github.com/benschw/cli-unit/files/cli-unit > $BIN_PATH
    chmod +x $BIN_PATH
fi

ARGS=( "$@" )

./$BIN_PATH ${ARGS[@]}
