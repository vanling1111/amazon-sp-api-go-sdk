.PHONY: help test lint fmt imports build clean install-tools coverage bench

# 默认目标
help:
	@echo "Amazon SP-API Go SDK - Makefile"
	@echo ""
	@echo "可用命令:"
	@echo "  make test          - 运行所有测试"
	@echo "  make test-unit     - 运行单元测试"
	@echo "  make test-integration - 运行集成测试"
	@echo "  make lint          - 运行代码检查"
	@echo "  make fmt           - 格式化代码"
	@echo "  make imports       - 整理 imports"
	@echo "  make build         - 构建项目"
	@echo "  make clean         - 清理构建产物"
	@echo "  make coverage      - 生成测试覆盖率报告"
	@echo "  make bench         - 运行基准测试"
	@echo "  make install-tools - 安装开发工具"

# 运行所有测试
test:
	go test ./... -v -race

# 运行单元测试
test-unit:
	go test ./... -v -race -short

# 运行集成测试
test-integration:
	go test ./tests/integration/... -v -race

# 代码检查
lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Run 'make install-tools'"; \
		exit 1; \
	fi

# 格式化代码
fmt:
	gofmt -s -w .
	go fmt ./...

# 整理 imports
imports:
	@if command -v goimports >/dev/null 2>&1; then \
		goimports -w .; \
	else \
		echo "goimports not installed. Run 'make install-tools'"; \
		exit 1; \
	fi

# 构建项目
build:
	go build ./...

# 清理构建产物
clean:
	go clean
	rm -f *.out
	rm -f *.test
	rm -rf coverage/

# 测试覆盖率
coverage:
	@mkdir -p coverage
	go test ./... -v -race -coverprofile=coverage/coverage.out -covermode=atomic
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@echo "Coverage report generated: coverage/coverage.html"

# 基准测试
bench:
	go test ./... -bench=. -benchmem

# 安装开发工具
install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golang/mock/mockgen@latest
	@echo "Tools installed successfully!"

# CI 流程
ci: fmt imports lint test coverage
	@echo "CI checks passed!"

