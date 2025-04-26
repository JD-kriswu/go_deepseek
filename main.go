package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"

	"github.com/kriswu/go_deepseek/grpc"
	"github.com/kriswu/go_deepseek/proto"
	"github.com/kriswu/go_deepseek/siliconproxy"
	grpclib "google.golang.org/grpc"
)

// 服务器配置结构体
type ServerConfig struct {
	URL      string `json:"url"`
	Token    string `json:"token"`
	GRPCPort int    `json:"grpc_port"`
}

// 加载服务器配置
func loadServerConfig(configPath string) (*ServerConfig, error) {
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return nil, fmt.Errorf("获取配置文件绝对路径失败: %w", err)
	}

	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var config ServerConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return &config, nil
}

func main() {
	// 加载服务器配置
	config, err := loadServerConfig("conf/server.conf")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 创建SiliconProxy实例
	sp := siliconproxy.NewSiliconProxy(config.Token)

	// 创建gRPC服务器
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRPCPort))
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	s := grpclib.NewServer()

	// 注册服务
	proto.RegisterSiliconServiceServer(s, grpc.NewSiliconServer(sp))

	// 启动服务
	log.Printf("gRPC服务器启动，监听端口: %d\n", config.GRPCPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
