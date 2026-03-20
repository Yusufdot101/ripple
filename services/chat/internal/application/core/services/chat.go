package services

import "github.com/Yusufdot101/ribble/services/chat/internal/ports"

type ChatService struct {
	repo ports.Repository
}

func NewChatService(repo ports.Repository) *ChatService {
	return &ChatService{
		repo: repo,
	}
}
