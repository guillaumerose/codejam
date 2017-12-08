#!/usr/bin/env bash
cat $1 | cut -d' ' -f1 | sed 's/$/ = 0/'
cat $1 | sed 's/inc/+=/' | sed 's/dec/-=/'
FOO=$(cat $1 | cut -d' ' -f1 | sed 's/$/,/' | tr -d '\n')
echo \[$FOO\].max