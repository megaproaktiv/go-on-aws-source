#!/bin/bash
# Call all go unit test in all exercise directories
# All task have its own directory code-task$n
basebase=`pwd`
for i in exercise*
do
  cd $basebase
  cd $i
  echo $i
  base=`pwd`
  for n in 1 2 3 4
  do
    if [ -d "code-task$n" ]; then
      echo "code-task$n"
      cd code-task$n
      go mod tidy
      go test -short ./...
      cd ..
    fi
    cd $base
  done
done
