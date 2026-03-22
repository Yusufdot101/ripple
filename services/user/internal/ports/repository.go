package ports

import "github.com/Yusufdot101/ribble/services/user/internal/application/core/domain"

type Repository interface {
	InsertUser(user *domain.User) error
	FindUserByProviderAndSub(provider, sub string) (*domain.User, error)

	InsertToken(token *domain.Token) error
	GetTokenByStringAndUse(tokenString string, tokenUse domain.TokenUse) (*domain.Token, error)
	DeleteTokenByStringAndUse(tokenString string, tokenUse domain.TokenUse) error
}
