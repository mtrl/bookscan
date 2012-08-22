#!/bin/bash
echo "Starting camera left"
nohup ./capture left $1 $2 $5&
echo "Starting camera right"
#nohup ./capture right $3 $4 $5&
tail -f nohup.out
