#!/bin/bash
rm nohup.out
echo "Starting camera left"
nohup ./capture left $1 $2 $5 $6&
echo "Starting camera right"
nohup ./capture right $3 $4 $5 $6&
tail -f nohup.out
