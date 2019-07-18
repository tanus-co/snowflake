#!/bin/bash

# build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o snowflake

# scp
scp -P 10017 snowflake Dockerfile docker-compose.yml root@117.187.225.202:/root/tanus
scp -P 10017 server-app.yml root@117.187.225.202:/root/tanus/app.yml
scp -P 10017 snowflake.conf root@117.187.225.202:/data/docker/nginx/conf.d/

# deploy
ssh root@117.187.225.202 -p 10017 "cd /root/tanus; docker-compose down; docker rmi snowflake:1.0.0; docker build -t snowflake:1.0.0 .; docker-compose up -d; rm -f snowflake Dockerfile *.yml; docker restart tanus_nginx_1"

# clean
rm -f snowflake
