package ports

import "github.com/Yusufdot101/ripple/services/chat/internal/application/core/domain"

type ChatService interface {
	NewChatWithParticipants(userIDs []uint) (uint, error)
	GetChatByUserIDs(userIDs []uint) (*domain.Chat, error)
}
