package siliconproxy

import (
	"encoding/json"
	"fmt"

	"github.com/kriswu/go_deepseek/httpclient"
)

const (
	// API基础URL
	BaseURL = "https://api.siliconflow.cn/v1"
	
	// API路径
	ChatCompletionsPath = "/chat/completions"
	EmbeddingsPath      = "/embeddings"
	RerankPath          = "/rerank"
	ImagesGenerationPath = "/images/generations"
	UploadVoicePath     = "/audio/voice"
	CreateSpeechPath    = "/audio/speech"
	VoiceListPath       = "/audio/voice/list"
	DeleteVoicePath     = "/audio/voice/delete"
	VideosSubmitPath    = "/videos/submit"
	VideosStatusPath    = "/videos/status"
	ModelsPath          = "/models"
	UserInfoPath        = "/userinfo"
)

// SiliconProxy 是Silicon Flow API的代理
type SiliconProxy struct {
	client *httpclient.Client
	token  string
}

// NewSiliconProxy 创建一个新的Silicon Flow API代理
func NewSiliconProxy(token string, options ...httpclient.ClientOption) *SiliconProxy {
	return &SiliconProxy{
		client: httpclient.NewClient(options...),
		token:  token,
	}
}

// 错误响应结构
type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Param   string `json:"param,omitempty"`
		Code    string `json:"code,omitempty"`
	} `json:"error"`
}

// 处理API响应，检查错误
func (sp *SiliconProxy) handleAPIResponse(resp *httpclient.Response, err error) (*httpclient.Response, error) {
	if err != nil {
		return nil, err
	}

	// 检查状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errResp ErrorResponse
		if err := json.Unmarshal([]byte(resp.Body), &errResp); err == nil {
			return nil, fmt.Errorf("API错误: %s (类型: %s, 代码: %s)", 
				errResp.Error.Message, errResp.Error.Type, errResp.Error.Code)
		}
		return nil, fmt.Errorf("API请求失败，状态码: %d, 响应: %s", resp.StatusCode, resp.Body)
	}

	return resp, nil
}