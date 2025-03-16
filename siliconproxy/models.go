package siliconproxy

import (
	"encoding/json"
	"fmt"
)

// ModelListResponse 表示模型列表响应
type ModelListResponse struct {
	Object string      `json:"object"`
	Data   []ModelData `json:"data"`
}

// ModelData 表示模型数据
type ModelData struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	OwnedBy string `json:"owned_by"`
}

// GetModelList 获取模型列表
func (sp *SiliconProxy) GetModelList() (*ModelListResponse, error) {
	url := BaseURL + ModelsPath
	
	// 发送GET请求
	resp, err := sp.client.GetWithAuth(url, sp.token)
	resp, err = sp.handleAPIResponse(resp, err)
	if err != nil {
		return nil, err
	}
	
	// 解析响应
	var result ModelListResponse
	if err := json.Unmarshal([]byte(resp.Body), &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	
	return &result, nil
}