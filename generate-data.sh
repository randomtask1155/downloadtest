#!/bin/bash


mkdir downloads

for i in 1 2 5 10 15 20 50 100 1024
do
	dd if=/dev/zero of=downloads/${i}mb bs=1048576 count=$i
done
