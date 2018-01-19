#!/bin/bash

echo -e "\e[1;32m step1: build ......"
echo -e "\e[0m"

#CGO_ENABLED=0 是一个编译标志，会让构建系统忽略cgo并且静态链接所有依赖；
#-a会强制重新编译，即使所有包都是由最新代码编译的；
#-installsuffix cgo 会为新编译的包目录添加一个后缀，这样可以把编译的输出与默认的路径分离。

CGO_ENABLED=0 go build -a -installsuffix cgo client.go

IMAGE_NAME="solinx.co/market/demo-client:0.1"

docker build -t $IMAGE_NAME .

if [ $? -ne 0 ]; then
  echo -e "\e[1;31m build docker is unsuccessful."
  echo -e "\e[0m"
  exit;
else
  echo -e "\e[1:31m summary: build docker image successful, image name :$IMAGE_NAME"
  echo -e "\e[0m"
fi