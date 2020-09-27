#!/bin/bash
docker build -t git-auto-sync .
docker rm -f git-auto-sync
docker run  --name git-auto-sync -v ~/.ssh:/root/.ssh \
-v ~/tmp/test:/src \
--privileged --pid=host \
git-auto-sync \
-p /src \
-v 4 \
--commit-interval 10s \
--name shenkonghui \
--email shenkh1992@gmail.com \
--push-interval 1m 
