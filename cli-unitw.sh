#!/bin/bash


BIN_PATH=".cli-unit/cli-unit.sh"

if [ ! -f $BIN_PATH ]; then
	mkdir -p $(dirname $BIN_PATH)
    curl -sS https://raw.githubusercontent.com/benschw/cli-unit/master/cli-unit.sh > $BIN_PATH
fi

ARGS=( "$@" )

/bin/bash $BIN_PATH ${ARGS[@]}
