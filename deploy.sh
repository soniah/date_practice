#!/bin/bash

dt=`date +%F_%T` # oh the irony
target="/var/tmp/data_practice_${dt}"

# backup; I can't be bothered doing git surgery to fix mistakes
cp -a /Users/sonia/go/src/github.com/soniah/date_practice $target

git co answers

ssed -i '/REPLACE_EMPTY_STRING/ s/\(\s\+\).*/\1"",/' dates_test.go
ssed -i '/REPLACE_NIL/ s/\(\s\+\).*/\1nil,/' dates_test.go
ssed -i '/TRUNCATE/,$ d' dates_test.go

cp dates_test.go /var/tmp
git stash
git checkout master
cp /var/tmp/dates_test.go .
