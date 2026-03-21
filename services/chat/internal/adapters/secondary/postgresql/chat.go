package postgresql

import (
	"github.com/Yusufdot101/ribble/services/chat/internal/application/core/domain"
	"github.com/Yusufdot101/ribble/services/chat/internal/ports"
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Participants []ChatParticipant `gorm:"constraint:OnDelete:CASCADE;"`
	Messages     []Message         `gorm:"constraint:OnDelete:CASCADE;"`
}

func (a *Adapter) InsertChat(chat *domain.Chat) error {
	chatModel := &Chat{}

	res := a.db.Save(chatModel)
	if res.Error == nil {
		chat.ID = chatModel.ID
	}
	return res.Error
}

func (a *Adapter) WithTx(fn func(repo ports.Repository) error) error {
	return a.db.Transaction(func(tx *gorm.DB) error {
		txRepo := &Adapter{db: tx} // same struct, but with tx as the db
		return fn(txRepo)
	})
}
