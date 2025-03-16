# Go项目Makefile

# 变量定义
BINARY_NAME=go_deepseek
GO=go
GOFLAGS=-v

# 默认目标
.PHONY: all
all: build

# 构建应用
.PHONY: build
build:
	$(GO) build $(GOFLAGS) -o $(BINARY_NAME) .

# 运行应用
.PHONY: run
run: build
	./$(BINARY_NAME)

# 清理生成的文件
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

# 初始化Go模块
.PHONY: init
init:
	$(GO) mod init github.com/kriswu/go_deepseek
	$(GO) mod tidy

# 测试
.PHONY: test
test:
	$(GO) test ./...

# 帮助信息
.PHONY: help
help:
	@echo "可用的命令:"
	@echo "  make build  - 编译应用"
	@echo "  make run    - 运行应用"
	@echo "  make clean  - 清理生成的文件"
	@echo "  make init   - 初始化Go模块"
	@echo "  make test   - 运行测试"
	@echo "  make help   - 显示帮助信息"