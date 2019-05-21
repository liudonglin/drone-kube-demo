FROM ubuntu:latest
#FROM scratch

#设置工作目录
WORKDIR /app

#将服务器的go工程代码加入到docker容器中
ADD ci-test /app

#暴露端口
EXPOSE 8000
#最终运行docker的命令
ENTRYPOINT  ["./ci-test"]

# 执行 docker build --rm=true -t xiayangqiushi/drone-ci-golang:1.0.0 .
# 执行 docker run -d -p 8000:8000 xiayangqiushi/drone-ci-golang:1.0.0