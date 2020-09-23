FROM golang:1.14

WORKDIR /src/git-auto-sync
COPY .  .
ENV GO111MODULE=on
ENV GOPROXY https://goproxy.io
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor


FROM alpine:3.10
WORKDIR /target
# 拷贝后端二进制文件
COPY --from=0 /src/git-auto-sync/git-auto-sync .

ENTRYPOINT [ "/target/git-auto-sync"]