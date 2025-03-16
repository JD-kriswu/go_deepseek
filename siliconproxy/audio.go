package siliconproxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// UploadVoiceRequest 表示上传参考音频请求
type UploadVoiceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	FilePath    string `json:"-"` // 不序列化到JSON
}

// UploadVoiceResponse 表示上传参考音频响应
type UploadVoiceResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

// CreateSpeechRequest 表示文本转语音请求
type CreateSpeechRequest struct {
	Model          string  `json:"model"`
	Input          string  `json:"input"`
	Voice          string  `json:"voice"`
	ResponseFormat string  `json:"response_format,omitempty"`
	Speed          float64 `json:"speed,omitempty"`
}

// VoiceListResponse 表示参考音频列表响应
type VoiceListResponse struct {
	Object string      `json:"object"`
	Data   []VoiceData `json:"data"`
}

// VoiceData 表示参考音频数据
type VoiceData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

// DeleteVoiceRequest 表示删除参考音频请求
type DeleteVoiceRequest struct {
	ID string `json:"id"`
}

// DeleteVoiceResponse 表示删除参考音频响应
type DeleteVoiceResponse struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

// UploadVoice 上传参考音频
func (sp *SiliconProxy) UploadVoice(req *UploadVoiceRequest) (*UploadVoiceResponse, error) {
	url := BaseURL + UploadVoicePath

	// 打开文件
	file, err := os.Open(req.FilePath)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %w", err)
	}
	defer file.Close()

	// 创建multipart表单
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加文件
	part, err := writer.CreateFormFile("file", filepath.Base(req.FilePath))
	if err != nil {
		return nil, fmt.Errorf("创建表单文件失败: %w", err)
	}

	if _, err = io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("复制文件内容失败: %w", err)
	}

	// 添加其他字段
	if err = writer.WriteField("name", req.Name); err != nil {
		return nil, fmt.Errorf("添加name字段失败: %w", err)
	}

	if req.Description != "" {
		if err = writer.WriteField("description", req.Description); err != nil {
			return nil, fmt.Errorf("添加description字段失败: %w", err)
		}
	}

	if err = writer.Close(); err != nil {
		return nil, fmt.Errorf("关闭writer失败: %w", err)
	}

	// 创建请求
	httpReq, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", sp.token))

	// 发送请求
	resp, err := sp.client.Post(url, writer.FormDataContentType(), body.String())
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	// 读取响应体
	respBody := []byte(resp.Body)

	// 检查状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errResp ErrorResponse
		if err := json.Unmarshal(respBody, &errResp); err == nil {
			return nil, fmt.Errorf("API错误: %s (类型: %s, 代码: %s)",
				errResp.Error.Message, errResp.Error.Type, errResp.Error.Code)
		}
		return nil, fmt.Errorf("API请求失败，状态码: %d, 响应: %s", resp.StatusCode, string(respBody))
	}

	// 解析响应
	var result UploadVoiceResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// CreateSpeech 创建文本转语音请求
func (sp *SiliconProxy) CreateSpeech(req *CreateSpeechRequest) ([]byte, error) {
	url := BaseURL + CreateSpeechPath

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

	// 返回原始响应体（音频数据）
	return []byte(resp.Body), nil
}

// GetVoiceList 获取参考音频列表
func (sp *SiliconProxy) GetVoiceList() (*VoiceListResponse, error) {
	url := BaseURL + VoiceListPath

	// 发送GET请求
	resp, err := sp.client.GetWithAuth(url, sp.token)
	resp, err = sp.handleAPIResponse(resp, err)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var result VoiceListResponse
	if err := json.Unmarshal([]byte(resp.Body), &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// DeleteVoice 删除参考音频
func (sp *SiliconProxy) DeleteVoice(req *DeleteVoiceRequest) (*DeleteVoiceResponse, error) {
	url := BaseURL + DeleteVoicePath

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
	var result DeleteVoiceResponse
	if err := json.Unmarshal([]byte(resp.Body), &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}
