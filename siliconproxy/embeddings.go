package siliconproxy

import (
	"encoding/json"
	"fmt"
)

// EmbeddingRequest 表示嵌入请求
type EmbeddingRequest struct {
	Model string   `json:"model"`
	Input interface{} `json:"input"`
	User  string   `json:"user,omitempty"`
}

// EmbeddingResponse 表示嵌入响应
type EmbeddingResponse struct {
	Object string             `json:"object"`
	Data   []EmbeddingData    `json:"data"`
	Model  string             `json:"model"`
	Usage  EmbeddingUsage     `json:"usage"`
}

// EmbeddingData 表示嵌入数据
type EmbeddingData struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

// EmbeddingUsage 表示嵌入使用情况
type EmbeddingUsage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

// CreateEmbedding 创建嵌入请求
func (sp *SiliconProxy) CreateEmbedding(req *EmbeddingRequest) (*EmbeddingResponse, error) {
	url := BaseURL + EmbeddingsPath
	
	// 将请求转换为JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("序列化请求失败: %w", err)
	}
	
	// 发送POST请求
	resp, err := sp.client.PostWithAuth(url, "application/json", string(reqBody), sp.token)
	resp, err = sp.handleAPIResponse(resp, err)
	if err != nil {
		return nil, err
	}
	
	// 解析响应
	var result EmbeddingResponse
	if err := json.Unmarshal([]byte(resp.Body), &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	
	return &result, nil
}