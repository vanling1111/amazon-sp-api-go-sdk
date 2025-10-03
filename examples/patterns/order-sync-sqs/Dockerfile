# 构建阶段
FROM golang:1.25-alpine AS builder

WORKDIR /build

# 复制 go.mod 和 go.sum
COPY ../../../go.mod ../../../go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY ../../../ .

# 构建二进制
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo \
    -o order-sync ./examples/patterns/order-sync-sqs/

# 运行阶段
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

WORKDIR /app

# 从构建阶段复制二进制
COPY --from=builder /build/order-sync .

# 创建日志目录
RUN mkdir -p /var/log

# 健康检查端点
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1

# 暴露端口
EXPOSE 8080 9090

# 运行
CMD ["./order-sync"]

