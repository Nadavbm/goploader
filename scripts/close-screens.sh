#!/bin/sh
SCREENS=$(screen -list | grep devses | awk '{print $1}')

for i in $SCREENS; do
   echo "$i"
   screen -XS $i quit
done
