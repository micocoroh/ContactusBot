package service

import (
	"context"

	"github.com/micocoroh/contactus/api"
)

type chatService struct{}

func NewChatbotService() api.ChatbotServiceServer {
	return &chatService{}
}

func (s *chatService) ListChatbot(context.Context, *api.ListChatbotRequest) (*api.ListChatbotResponse, error) {
	return nil, nil
}
