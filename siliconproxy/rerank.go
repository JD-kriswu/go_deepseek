package siliconproxy

import (
	"encoding/json"
	"fmt"
)

// RerankRequest 表示重排序请求
type RerankRequest struct {
	Model    string   `json:"model"`
	Query    string   `json:"query"`
	Documents []string `json:"documents"`
	TopN     int      `json:"top_n,omitempty"`
	User     string   `json:"user,omitempty"`
}

// RerankResponse 表示重排序响应
type RerankResponse struct {
	Object string        `json:"object"`
	Model  string        `json:"model"`
	Results []RerankResult `json:"results"`
	Usage  RerankUsage     `json:"usage"`
}

// RerankResult 表示重排序结果
type RerankResult struct {
	Index       int     `json:"index"`
	Document    string  `json:"document"`
	RelevanceScore float64 `json:"relevance_score"`
}

// RerankUsage 表示重排序使用情况
type RerankUsage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

// CreateRerank 创建重排序请求
func (sp *SiliconProxy) CreateRerank(req *RerankRequest) (*RerankResponse, error) {
	url := BaseURL + RerankPath
	
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
	var result RerankResponse
	if err := json.Unmarshal([]byte(resp.Body), &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	
	return &result, nil
}