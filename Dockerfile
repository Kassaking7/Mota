# 使用官方的Golang基础镜像
FROM golang:1.20

# 设置工作目录
WORKDIR /app

# 将go.mod和go.sum复制到工作目录中并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 将项目的所有文件复制到工作目录中
COPY . .

# 构建Go应用
RUN go build -o main

# 暴露端口
EXPOSE 8088

# 启动应用
CMD ["./main"]
