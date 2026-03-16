package postgresql

import (
	"context"
	"time"

	"github.com/Yusufdot101/ribble/services/user/internal/application/core/domain"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	TokenString string
	UserID      uint
	Expires     time.Time
	Use         domain.TokenUse
	TokenType   domain.TokenType
}

// InsertToken(token *domain.Token) error
// GetTokenByStringAndUse(tokenString string, tokenUse domain.TokenUse) (*domain.Token, error)
// DeleteTokenByStringAndUse(tokenString string, tokenUse domain.TokenUse) error

func (a *Adapter) InsertToken(token *domain.Token) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tokenModel := &Token{
		TokenString: token.TokenString,
		UserID:      token.UserID,
		Use:         token.Use,
		Expires:     token.Expires,
		TokenType:   token.TokenType,
	}
	res := a.DB.WithContext(ctx).Create(tokenModel)
	if res.Error == nil {
		token.ID = tokenModel.ID
	}

	return res.Error
}
