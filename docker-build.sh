#! /bin/bash
version=$1
push=$2
if [ "$1" == "" ]; then
  version=latest
fi

env GOOS=linux GOARCH=amd64 go build -o app

docker build  -t docker.hunantv.com/huyinghuan/oauth:$version .

if [ "$version" != "latest" ]; then
    docker tag docker.hunantv.com/huyinghuan/oauth:$version docker.hunantv.com/huyinghuan/oauth:latest
    docker tag docker.hunantv.com/huyinghuan/oauth:$version huyinghuan/oauth:$version
fi

if [ "$2" != "" ]; then
  docker push docker.hunantv.com/huyinghuan/oauth:$version
fi

rm app