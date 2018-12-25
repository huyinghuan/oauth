#! /bin/bash
version=$1
if [ "$1" == "" ]; then
  version=latest
fi

docker build  -t docker.hunantv.com/huyinghuan/oauth:$version .

if [ "$version" != "latest" ]; then
    docker tag docker.hunantv.com/huyinghuan/oauth:$version docker.hunantv.com/huyinghuan/oauth:latest
    docker tag docker.hunantv.com/huyinghuan/oauth:$version huyinghuan/oauth:$version
fi