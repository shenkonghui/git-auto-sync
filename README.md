# git-auto-sync
Automatically synchronize files with git

Run with Binary
```
./git-auto-sync \
 -p ~/tmp/test \
 -v 4 \
 --commit-interval 10s \
 --name shenkonghui \
 --email shenkh1992@gmail.com \
 --push-interval 1m 
```

Run with Docker(failed)
```
docker run  --name git-auto-sync -v ~/.ssh:/root/.ssh \
-v ~/tmp/test:/src \
-v ${SSH_AUTH_SOCK}:${SSH_AUTH_SOCK} \
-e SSH_AUTH_SOCK="${SSH_AUTH_SOCK}" \
git-auto-sync \
-p /src \
-v 4 \
--commit-interval 10s \
--name shenkonghui \
--email shenkh1992@gmail.com \
--push-interval 1m 
```