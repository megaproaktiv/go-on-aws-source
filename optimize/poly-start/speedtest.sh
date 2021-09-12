#!/bin/bash
echo "Node.JS"
time node index.js
time node index.js
echo  "Python"
time python3 hello.py
time python3 hello.py
echo "Go run"
time go run main.go
echo "Build Go"
go build -o hello main.go
echo "Execute static binary"
time ./hello  
time ./hello  

