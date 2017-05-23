#!/bin/sh
echo Building pchico83/demo:build

docker build -t pchico83/demo:build -f Dockerfile.build .

docker create --name extract pchico83/demo:build  
docker cp extract:/go/src/github.com/pchico83/demo/app ./app  
docker rm -f extract

echo Building pchico83/demo:meetup

docker build -t pchico83/demo:meetup .
rm ./app
