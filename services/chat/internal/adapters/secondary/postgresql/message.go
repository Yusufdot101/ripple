package postgresql

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Yusufdot101/ripple/services/chat/internal/application/core/domain"
	"gorm.io/gorm"
)

type Message struct {
	ID          uint `gorm:"primarykey"`
	ChatID      uint
	SenderID    uint
	Content     string
	Deleted     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Status      domain.MessageStatus
	MessageType domain.MessageType
}

func (a *Adapter) InsertMessage(message *domain.Message) error {
	messageModel := &Message{
		ChatID:      message.ChatID,
		SenderID:    message.SenderID,
		Content:     message.Content,
		Status:      domain.MessageDelivered,
		MessageType: message.MessageType,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := a.db.WithContext(ctx).Save(messageModel)
	if res.Error == nil {
		message.ID = messageModel.ID
		message.CreatedAt = messageModel.CreatedAt
		message.UpdatedAt = messageModel.UpdatedAt
		message.Status = domain.MessageDelivered
	}
	return res.Error
}

func (a *Adapter) GetMessages(chatID uint, messageFilter domain.GetMessageFilter) ([]*domain.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var messageModels []*Message
	res := a.db.WithContext(ctx).
		Where("chat_id = ? AND id > ?", chatID, messageFilter.LastMessageID).
		Order("created_at ASC").
		Find(&messageModels)

	if res.Error != nil {
		return nil, res.Error
	}

	messages := []*domain.Message{}
	for _, messageModel := range messageModels {
		status := messageModel.Status
		if status == "" {
			status = domain.MessageDelivered
		}
		message := &domain.Message{
			ID:          messageModel.ID,
			ChatID:      messageModel.ChatID,
			CreatedAt:   messageModel.CreatedAt,
			UpdatedAt:   messageModel.UpdatedAt,
			SenderID:    messageModel.SenderID,
			Content:     messageModel.Content,
			DeletedAt:   messageModel.DeletedAt,
			Deleted:     messageModel.Deleted,
			Status:      status,
			MessageType: messageModel.MessageType,
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func (a *Adapter) DeleteMessage(chatID, userID, messageID uint) (*domain.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	messageModel := &Message{}

	err := a.db.WithContext(ctx).
		Where("id = ? AND sender_id = ?", messageID, userID).
		First(messageModel).Error
	if err != nil {
		return nil, err
	}

	content := "Message deleted"
	res := a.db.WithContext(ctx).
		Model(&Message{}).
		Where("id = ? AND sender_id = ? AND chat_id = ?", messageID, userID, chatID).
		Updates(map[string]any{
			"content":    content,
			"deleted_at": gorm.Expr("Now()"),
			"deleted":    true,
		})
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, domain.ErrRecordNotFound
	}

	status := messageModel.Status
	if status == "" {
		status = domain.MessageDelivered
	}
	message := &domain.Message{
		ID:          messageModel.ID,
		ChatID:      messageModel.ChatID,
		CreatedAt:   messageModel.CreatedAt,
		DeletedAt:   messageModel.DeletedAt,
		MessageType: messageModel.MessageType,
		SenderID:    messageModel.SenderID,
		Content:     content,
		UpdatedAt:   messageModel.UpdatedAt,
		Deleted:     messageModel.Deleted,
		Status:      status,
	}

	return message, nil
}

func (a *Adapter) DeleteAnyMessage(chatID, userID uint, username string, messageID uint) (*domain.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var messageModel Message

	err := a.db.WithContext(ctx).
		Where("id = ? AND chat_id = ?", messageID, chatID).
		First(&messageModel).Error
	if err != nil {
		return nil, err
	}

	content := "Message deleted"
	if messageModel.SenderID != userID {
		content = fmt.Sprintf("Message deleted by %s", username)
	}

	res := a.db.WithContext(ctx).
		Model(&Message{}).
		Where("id = ? AND chat_id = ?", messageID, chatID).
		Updates(map[string]any{
			"content":    content,
			"deleted_at": gorm.Expr("Now()"),
			"deleted":    true,
		})
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, domain.ErrRecordNotFound
	}

	status := messageModel.Status
	if status == "" {
		status = domain.MessageDelivered
	}
	message := &domain.Message{
		ID:          messageModel.ID,
		ChatID:      messageModel.ChatID,
		CreatedAt:   messageModel.CreatedAt,
		DeletedAt:   messageModel.DeletedAt,
		MessageType: messageModel.MessageType,
		SenderID:    messageModel.SenderID,
		Content:     content,
		UpdatedAt:   messageModel.UpdatedAt,
		Deleted:     messageModel.Deleted,
		Status:      status,
	}

	return message, nil
}

func (a *Adapter) EditMessage(userID, messageID uint, newContent string) (*domain.Message, error) {
	if strings.TrimSpace(newContent) == "" {
		return nil, domain.ErrInvalidMessageContent
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	messageModel := &Message{}
	// fetch the message
	err := a.db.WithContext(ctx).Where("id = ? AND sender_id = ?", messageID, userID).First(messageModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrRecordNotFound
		}
		return nil, err
	}

	// verify updateablity
	updateWindow := time.Hour
	if time.Since(messageModel.CreatedAt) > updateWindow {
		return nil, domain.ErrUpdateWindowOver
	}

	// update
	err = a.db.WithContext(ctx).
		Model(&Message{}).
		Where("id = ? AND sender_id = ?", messageID, userID).
		Update("content", newContent).Error
	if err != nil {
		return nil, err
	}

	// fetch the updated message
	err = a.db.WithContext(ctx).Where("id = ? AND sender_id = ?", messageID, userID).First(messageModel).Error
	if err != nil {
		return nil, err
	}

	message := &domain.Message{
		ID:          messageModel.ID,
		Content:     messageModel.Content,
		CreatedAt:   messageModel.CreatedAt,
		UpdatedAt:   messageModel.UpdatedAt,
		ChatID:      messageModel.ChatID,
		SenderID:    messageModel.SenderID,
		DeletedAt:   messageModel.DeletedAt,
		Deleted:     messageModel.Deleted,
		Status:      messageModel.Status,
		MessageType: messageModel.MessageType,
	}

	return message, err
}
