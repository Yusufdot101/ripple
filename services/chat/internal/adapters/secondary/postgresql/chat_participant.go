package postgresql

import (
	"context"
	"errors"
	"time"

	"github.com/Yusufdot101/ripple/services/chat/internal/application/core/domain"
	"gorm.io/gorm"
)

type ChatParticipant struct {
	gorm.Model
	UserID     uint `gorm:"uniqueIndex:user_chat_idx"`
	ChatID     uint `gorm:"uniqueIndex:user_chat_idx"`
	ChatRoleID uint
}

func (a *Adapter) InsertChatParticipant(chatParticipant *domain.ChatParticipant) error {
	chatParticipantModel := &ChatParticipant{
		UserID:     chatParticipant.UserID,
		ChatID:     chatParticipant.ChatID,
		ChatRoleID: chatParticipant.ChatRoleID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := a.db.WithContext(ctx).Save(chatParticipantModel)
	if res.Error == nil {
		chatParticipant.ID = chatParticipantModel.ID
	}

	return res.Error
}

func (a *Adapter) GetChatUsers(chatID, currentUserID uint) ([]*domain.ChatParticipant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// check user is in the chat
	err := a.db.WithContext(ctx).
		Where("chat_id = ? AND user_id = ?", chatID, currentUserID).
		Find(&ChatParticipant{}).
		Error
	if err != nil {
		return nil, errors.New("not in chat")
	}

	// get the chat users
	chatParticipantModels := []*ChatParticipant{}
	err = a.db.WithContext(ctx).
		Where("chat_participants.chat_id = ?", chatID).
		Find(&chatParticipantModels).Error
	if err != nil {
		return nil, err
	}

	chatParticipants := []*domain.ChatParticipant{}
	for _, chatParticipantModel := range chatParticipantModels {
		chatParticipant := &domain.ChatParticipant{
			ID:     chatParticipantModel.ID,
			UserID: chatParticipantModel.UserID,
			ChatID: chatParticipantModel.ChatID,
		}
		chatParticipants = append(chatParticipants, chatParticipant)
	}

	return chatParticipants, nil
}
