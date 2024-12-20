# 拉取golang:1.22.2轻量镜像
FROM golang:1.22.2-alpine3.19 AS builder

# 配置模块代理
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

## 在docker的根目录下创建相应的使用目录
RUN mkdir -p /www/webapp

## 设置工作目录
WORKDIR /www/webapp

## 将后端的gin代码复制到docker容器中
COPY . .

#go构建可执行文件
RUN go build main.go

RUN chmod +x main

FROM alpine:latest

COPY --from=builder /www/webapp /usr/local/bin

WORKDIR /usr/local/bin

EXPOSE 8081

CMD ["./main"]