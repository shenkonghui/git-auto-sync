FROM golang:1.14

WORKDIR /src/git-auto-sync
COPY .  .
ENV GO111MODULE=on
ENV GOPROXY https://goproxy.io
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor


FROM ubuntu:latest
RUN apt-get update && apt-get install -y openssh-client && ssh-add -l
WORKDIR /target
# 拷贝后端二进制文件
COPY --from=0 /src/git-auto-sync/git-auto-sync .
#RUN  ssh-add -l
ENTRYPOINT [ "/target/git-auto-sync"]