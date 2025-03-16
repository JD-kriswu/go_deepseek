package siliconproxy

import (
	"encoding/json"
	"fmt"
)

// UserInfoResponse 表示用户账户信息响应
type UserInfoResponse struct {
	Object  string `json:"object"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Created int64  `json:"created"`
	Balance int64  `json:"balance"`
}

// GetUserInfo 获取用户账户信息
func (sp *SiliconProxy) GetUserInfo() (*UserInfoResponse, error) {
	url := BaseURL + UserInfoPath
	
	// 发送GET请求
	resp, err := sp.client.GetWithAuth(url, sp.token)
	resp, err = sp.handleAPIResponse(resp, err)
	if err != nil {
		return nil, err
	}
	
	// 解析响应
	var result UserInfoResponse
	if err := json.Unmarshal([]byte(resp.Body), &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	
	return &result, nil
}