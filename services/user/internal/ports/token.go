package ports

import "github.com/Yusufdot101/ribble/services/user/internal/application/core/domain"

type TokenService interface {
	New(tokenType domain.TokenType, tokenUse domain.TokenUse, userID uint) (*domain.Token, error)
	Save(token *domain.Token) error
}
