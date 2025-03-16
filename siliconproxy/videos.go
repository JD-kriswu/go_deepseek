package siliconproxy

import (
	"encoding/json"
	"fmt"
)

// VideoSubmitRequest 表示视频生成请求
type VideoSubmitRequest struct {
	Model       string `json:"model"`
	Prompt      string `json:"prompt"`
	Duration    int    `json:"duration,omitempty"`
	Quality     string `json:"quality,omitempty"`
	Style       string `json:"style,omitempty"`
	User        string `json:"user,omitempty"`
}

// VideoSubmitResponse 表示视频生成响应
type VideoSubmitResponse struct {
	ID      string `json:"id"`
	Created int64  `json:"created"`
	Status  string `json:"status"`
}

// VideoStatusRequest 表示获取视频状态请求
type VideoStatusRequest struct {
	ID string `json:"id"`
}

// VideoStatusResponse 表示获取视频状态响应
type VideoStatusResponse struct {
	ID      string `json:"id"`
	Created int64  `json:"created"`
	Status  string `json:"status"`
	URL     string `json:"url,omitempty"`
}

// CreateVideoSubmit 创建视频生成请求
func (sp *SiliconProxy) CreateVideoSubmit(req *VideoSubmitRequest) (*VideoSubmitResponse, error) {
	url := BaseURL + VideosSubmitPath
	
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
	var result VideoSubmitResponse
	if err := json.Unmarshal([]byte(resp.Body), &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	
	return &result, nil
}

// GetVideoStatus 获取视频状态
func (sp *SiliconProxy) GetVideoStatus(req *VideoStatusRequest) (*VideoStatusResponse, error) {
	url := BaseURL + VideosStatusPath
	
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
	var result VideoStatusResponse
	if err := json.Unmarshal([]byte(resp.Body), &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	
	return &result, nil
}