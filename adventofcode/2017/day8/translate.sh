#!/usr/bin/env bash
FOO=$(cat $1 | cut -d' ' -f1 | sed 's/$/,/' | tr -d '\n')
echo max = 0
cat $1 | cut -d' ' -f1 | sed 's/$/ = 0/'
cat $1 | sed 's/inc/+=/' | sed 's/dec/-=/' | sed s/$/\;max=\[max,$FOO\].max/
echo max