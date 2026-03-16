package ports

import (
	"context"

	"github.com/Yusufdot101/ribble/services/user/internal/application/core/domain"
)

type AuthService interface {
	NewUser(user *domain.User) error
	BeginAuth() (authURL, state, nonce string)
	HandleCallback(ctx context.Context, code, state, nonce string) (refreshToken, accessToken string, err error)
}
