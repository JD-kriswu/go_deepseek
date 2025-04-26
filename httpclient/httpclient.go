package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Response 表示HTTP响应
type Response struct {
	StatusCode int
	Headers    http.Header
	Body       string
}

// Client 是HTTP客户端
type Client struct {
	httpClient *http.Client
	headers    map[string]string
}

// ClientOption 定义客户端选项函数类型
type ClientOption func(*Client)

// WithTimeout 设置超时时间
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithHeader 添加请求头
func WithHeader(key, value string) ClientOption {
	return func(c *Client) {
		c.headers[key] = value
	}
}

// NewClient 创建新的HTTP客户端
func NewClient(options ...ClientOption) *Client {
	client := &Client{
		httpClient: &http.Client{
			Timeout: 100 * time.Second,
		},
		headers: make(map[string]string),
	}

	// 应用选项
	for _, option := range options {
		option(client)
	}

	return client
}

// 设置请求头
func (c *Client) setHeaders(req *http.Request) {
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}
}

// 处理HTTP响应
func (c *Client) handleResponse(resp *http.Response) (*Response, error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %w", err)
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       string(body),
	}, nil
}

// Get 发送GET请求
func (c *Client) Get(url string) (*Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	c.setHeaders(req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}

	return c.handleResponse(resp)
}

// GetWithAuth 发送带有Authorization的GET请求
func (c *Client) GetWithAuth(url, token string) (*Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	c.setHeaders(req)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}

	return c.handleResponse(resp)
}

// Post 发送POST请求
func (c *Client) Post(url, contentType, body string) (*Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(body))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", contentType)
	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}

	return c.handleResponse(resp)
}

// PostWithAuth 发送带有Authorization的POST请求
func (c *Client) PostWithAuth(url, contentType, body, token string) (*Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(body))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}

	return c.handleResponse(resp)
}

// Put 发送PUT请求
func (c *Client) Put(url, contentType, body string) (*Response, error) {
	req, err := http.NewRequest("PUT", url, bytes.NewBufferString(body))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", contentType)
	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}

	return c.handleResponse(resp)
}

// Delete 发送DELETE请求
func (c *Client) Delete(url string) (*Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	c.setHeaders(req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}

	return c.handleResponse(resp)
}
