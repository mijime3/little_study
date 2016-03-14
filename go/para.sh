#!/bin/bash

echo `date`

a=0
if [ $# -ne 0 ]
then
  cnt=$1
else
  cnt=10
fi
while [ $a -ne $cnt ]
do
  num=`expr 30000 + 100 \* $a`
  a=`expr $a + 1`
  ./sum_prime -num=$num &
done
