package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/kriswu/go_deepseek/siliconproxy"
)

// 服务器配置结构体
type ServerConfig struct {
	URL   string `json:"url"`
	Token string `json:"token"`
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

	// 测试获取模型列表
	fmt.Println("测试获取模型列表:")
	models, err := sp.GetModelList()
	if err != nil {
		log.Printf("获取模型列表失败: %v\n", err)
	} else {
		fmt.Printf("获取到 %d 个模型\n", len(models.Data))
		for _, model := range models.Data {
			fmt.Printf("模型ID: %s, 所有者: %s\n", model.ID, model.OwnedBy)
		}
	}
	fmt.Println()

	// 测试对话能力
	fmt.Println("测试对话能力:")
	// fullfill all the fields of ChatCompletionRequest

	chatReq := &siliconproxy.ChatCompletionRequest{
		Model: "Qwen/QwQ-32B",
		Messages: []siliconproxy.ChatCompletionMessage{
			{Role: "user", Content: "What opportunities and challenges will the Chinese large model industry face in 2025?"},
		},
		Stream:           false,
		MaxTokens:        512,
		Stop:             nil,
		Temperature:      0.7,
		TopP:             0.7,
		TopK:             50,
		FrequencyPenalty: 0.5,
		N:                1,
		ResponseFormat: &siliconproxy.ResponseFormat{
			Type: "text",
		},
		Tools: []siliconproxy.Tool{
			{
				Type: "function",
				Function: siliconproxy.FunctionObject{
					Description: "test",
					Name:        "test",
					Parameters:  map[string]interface{}{},
					Strict:      false,
				},
			},
		},
	}
	chatResp, err := sp.CreateChatCompletion(chatReq)
	if err != nil {
		log.Printf("创建对话失败: %v\n", err)
	} else {
		fmt.Printf("对话内容: %s\n", chatResp.Choices[0].Message.Content)
		fmt.Printf("使用Token数: %d\n", chatResp.Usage.TotalTokens)
	}

	fmt.Println()
}
