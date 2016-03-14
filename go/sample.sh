#!/bin/bash

# -eq : equal
i=0
j=0
while [ $i -eq 0 ]
do
  j=`expr $j + 1`
  if [ $j -eq 5 ]; then
    i=`expr $i + 1`
  fi
  echo "(-eq : equal) i=$i, j=$j"
done
