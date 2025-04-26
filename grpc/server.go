package grpc

import (
	"context"
	"fmt"

	"github.com/kriswu/go_deepseek/proto"
	"github.com/kriswu/go_deepseek/siliconproxy"
)

// SiliconServer 实现gRPC服务接口
type SiliconServer struct {
	proto.UnimplementedSiliconServiceServer
	sp *siliconproxy.SiliconProxy
}

// NewSiliconServer 创建新的服务实例
func NewSiliconServer(sp *siliconproxy.SiliconProxy) *SiliconServer {
	return &SiliconServer{sp: sp}
}

// GetModelList 获取模型列表
func (s *SiliconServer) GetModelList(ctx context.Context, _ *proto.Empty) (*proto.GetModelListResponse, error) {
	models, err := s.sp.GetModelList()
	if err != nil {
		return nil, fmt.Errorf("获取模型列表失败: %w", err)
	}

	response := &proto.GetModelListResponse{}
	for _, model := range models.Data {
		response.Data = append(response.Data, &proto.Model{
			Id:      model.ID,
			OwnedBy: model.OwnedBy,
		})
	}

	return response, nil
}

// CreateChatCompletion 创建聊天对话
func (s *SiliconServer) CreateChatCompletion(ctx context.Context, req *proto.ChatCompletionRequest) (*proto.ChatCompletionResponse, error) {
	// 转换请求
	messages := make([]siliconproxy.ChatCompletionMessage, len(req.Messages))
	for i, msg := range req.Messages {
		messages[i] = siliconproxy.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	tools := make([]siliconproxy.Tool, len(req.Tools))
	for i, tool := range req.Tools {
		tools[i] = siliconproxy.Tool{
			Type: tool.Type,
			Function: siliconproxy.FunctionObject{
				Description: tool.Function.Description,
				Name:        tool.Function.Name,
				Strict:      tool.Function.Strict,
			},
		}
	}

	chatReq := &siliconproxy.ChatCompletionRequest{
		Model:            req.Model,
		Messages:         messages,
		Stream:           req.Stream,
		MaxTokens:        int(req.MaxTokens),
		Stop:             req.Stop,
		Temperature:      float64(req.Temperature),
		TopP:             float64(req.TopP),
		TopK:             int(req.TopK),
		FrequencyPenalty: float64(req.FrequencyPenalty),
		N:                int(req.N),
		ResponseFormat: &siliconproxy.ResponseFormat{
			Type: req.ResponseFormat.Type,
		},
		Tools: tools,
	}

	// 调用API
	chatResp, err := s.sp.CreateChatCompletion(chatReq)
	if err != nil {
		return nil, fmt.Errorf("创建对话失败: %w", err)
	}

	// 转换响应
	choices := make([]*proto.Choice, len(chatResp.Choices))
	for i, choice := range chatResp.Choices {
		choices[i] = &proto.Choice{
			Message: &proto.ChatMessage{
				Role:    choice.Message.Role,
				Content: choice.Message.Content,
			},
			Index: int32(i),
		}
	}

	response := &proto.ChatCompletionResponse{
		Id:      chatResp.ID,
		Object:  chatResp.Object,
		Created: chatResp.Created,
		Model:   chatResp.Model,
		Choices: choices,
		Usage: &proto.Usage{
			PromptTokens:     int32(chatResp.Usage.PromptTokens),
			CompletionTokens: int32(chatResp.Usage.CompletionTokens),
			TotalTokens:      int32(chatResp.Usage.TotalTokens),
		},
	}

	return response, nil
}
