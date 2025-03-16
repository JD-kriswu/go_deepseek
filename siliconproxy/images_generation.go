package siliconproxy

import (
	"encoding/json"
	"fmt"
)

// ImageGenerationRequest 表示图像生成请求
type ImageGenerationRequest struct {
	Model          string  `json:"model"`
	Prompt         string  `json:"prompt"`
	N              int     `json:"n,omitempty"`
	Size           string  `json:"size,omitempty"`
	ResponseFormat string  `json:"response_format,omitempty"`
	User           string  `json:"user,omitempty"`
}

// ImageGenerationResponse 表示图像生成响应
type ImageGenerationResponse struct {
	Created int64                `json:"created"`
	Data    []ImageGenerationData `json:"data"`
}

// ImageGenerationData 表示图像生成数据
type ImageGenerationData struct {
	URL     string `json:"url,omitempty"`
	B64JSON string `json:"b64_json,omitempty"`
}

// CreateImageGeneration 创建图像生成请求
func (sp *SiliconProxy) CreateImageGeneration(req *ImageGenerationRequest) (*ImageGenerationResponse, error) {
	url := BaseURL + ImagesGenerationPath
	
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
	var result ImageGenerationResponse
	if err := json.Unmarshal([]byte(resp.Body), &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	
	return &result, nil
}