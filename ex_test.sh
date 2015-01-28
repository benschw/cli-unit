#!/bin/bash

#test -json flag should make output be in json
#shell
cat ./ex.json | ./jsonfilter -json myArray

#output
["foo","bar"]


#test should target array element
#shell
cat ./ex.json | ./jsonfilter myArray.0

#output
foo
