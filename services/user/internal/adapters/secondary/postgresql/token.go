package postgresql

import (
	"context"
	"errors"
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

func (a *Adapter) GetTokenByStringAndUse(tokenString string, tokenUse domain.TokenUse) (
	*domain.Token, error,
) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tokenModel := &Token{}
	res := a.DB.WithContext(ctx).First(tokenModel, "expires > ?", time.Now())
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrRecordNotFound
		}
		return nil, res.Error
	}
	token := domain.NewToken(tokenModel.Use, tokenModel.TokenType, tokenModel.UserID, tokenModel.TokenString, tokenModel.Expires)
	return token, nil
}
