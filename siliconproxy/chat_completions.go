package siliconproxy

import (
	"encoding/json"
	"fmt"
)

// ChatCompletionRequest 表示聊天完成请求
type ChatCompletionRequest struct {
	Model            string                  `json:"model"`
	Messages         []ChatCompletionMessage `json:"messages"`
	Stream           bool                    `json:"stream,omitempty"`
	MaxTokens        int                     `json:"max_tokens,omitempty"`
	Stop             interface{}             `json:"stop,omitempty"`
	Temperature      float64                 `json:"temperature,omitempty"`
	TopP             float64                 `json:"top_p,omitempty"`
	TopK             int                     `json:"top_k,omitempty"`
	FrequencyPenalty float64                 `json:"frequency_penalty,omitempty"`
	N                int                     `json:"n,omitempty"`
	ResponseFormat   *ResponseFormat         `json:"response_format,omitempty"`
	Tools            []Tool                  `json:"tools,omitempty"`
}

// ChatCompletionMessage 表示聊天消息
type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ResponseFormat 表示响应格式
type ResponseFormat struct {
	Type string `json:"type"`
}

// Tool 表示工具
type Tool struct {
	Type     string         `json:"type"`
	Function FunctionObject `json:"function"`
}

// FunctionObject 表示函数对象
type FunctionObject struct {
	Description string                 `json:"description,omitempty"`
	Name        string                 `json:"name"`
	Parameters  map[string]interface{} `json:"parameters,omitempty"`
	Strict      bool                   `json:"strict,omitempty"`
}

// ChatCompletionResponse 表示聊天完成响应
type ChatCompletionResponse struct {
	ID      string                   `json:"id"`
	Choices []ChatCompletionChoice   `json:"choices"`
	Usage   ChatCompletionUsage      `json:"usage"`
	Created int64                    `json:"created"`
	Model   string                   `json:"model"`
	Object  string                   `json:"object"`
}

// ChatCompletionChoice 表示聊天完成选择
type ChatCompletionChoice struct {
	Message      ChatCompletionResponseMessage `json:"message"`
	FinishReason string                        `json:"finish_reason"`
}

// ChatCompletionResponseMessage 表示聊天完成响应消息
type ChatCompletionResponseMessage struct {
	Role             string     `json:"role"`
	Content          string     `json:"content"`
	ReasoningContent string     `json:"reasoning_content,omitempty"`
	ToolCalls        []ToolCall `json:"tool_calls,omitempty"`
}

// ToolCall 表示工具调用
type ToolCall struct {
	ID       string           `json:"id"`
	Type     string           `json:"type"`
	Function FunctionCallInfo `json:"function"`
}

// FunctionCallInfo 表示函数调用信息
type FunctionCallInfo struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

// ChatCompletionUsage 表示聊天完成使用情况
type ChatCompletionUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// CreateChatCompletion 创建聊天完成请求
func (sp *SiliconProxy) CreateChatCompletion(req *ChatCompletionRequest) (*ChatCompletionResponse, error) {
	url := BaseURL + ChatCompletionsPath
	
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
	var result ChatCompletionResponse
	if err := json.Unmarshal([]byte(resp.Body), &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	
	return &result, nil
}

// CreateChatCompletionStream 创建流式聊天完成请求
// 注意：此方法返回原始响应，客户端需要自行处理SSE流
func (sp *SiliconProxy) CreateChatCompletionStream(req *ChatCompletionRequest) (*ChatCompletionResponse, error) {
	// 确保流式标志设置为true
	req.Stream = true
	
	// 使用普通的CreateChatCompletion方法
	// 注意：实际实现中，应该处理SSE流，这里简化处理
	return sp.CreateChatCompletion(req)
}