package services

import (
	"log"

	"github.com/Yusufdot101/ribble/services/chat/internal/application/core/domain"
	"github.com/Yusufdot101/ribble/services/chat/internal/ports"
)

type ChatService struct {
	repo ports.Repository
}

func NewChatService(repo ports.Repository) *ChatService {
	return &ChatService{
		repo: repo,
	}
}

func (csvc *ChatService) NewChatWithParticipants(userIDs []uint) (uint, error) {
	log.Println("start", userIDs)
	chatID := uint(0)
	err := csvc.repo.WithTx(func(repo ports.Repository) error {
		log.Println("beginning")
		chat := domain.NewChat()
		err := repo.InsertChat(chat)
		if err != nil {
			return err
		}
		log.Println("chat: ", chat)
		chatID = chat.ID

		for _, userID := range userIDs {
			participant := domain.NewChatParticipant(userID, chat.ID)
			err = repo.InsertChatParticipant(participant)
			if err != nil {
				return err
			}
			log.Println("participant: ", participant)
		}
		log.Println("ending")
		return nil
	})
	if err != nil {
		return 0, err
	}

	return chatID, nil
}
