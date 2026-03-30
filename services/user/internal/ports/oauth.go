package ports

import (
	"context"

	"github.com/Yusufdot101/ripple/services/user/internal/application/core/domain"
)

type OAuthProvider interface {
	GetAuthURL(state, nonce string) string
	GetUserInfo(ctx context.Context, code, nonce string) (*domain.User, error)
}
