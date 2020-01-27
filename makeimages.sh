#!/bin/bash

cp ~/.macaddress macaddress

docker build -t python_test . --no-cache -f Dockerfile_python
docker build -t shell_test . --no-cache -f Dockerfile_shell
docker build -t go_test1 . --no-cache -f Dockerfile_go1
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
docker build -t go_test2 . --no-cache -f Dockerfile_go2
rm main
rm macaddress

