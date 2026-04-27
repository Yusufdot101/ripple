package ports

import (
	"context"

	"github.com/Yusufdot101/ripple/services/user/internal/application/core/domain"
)

type UserService interface {
	GetUsersByEmail(email string) ([]*domain.User, error)
	GetUsersByIDs(ctx context.Context, userIDs []uint32) ([]*domain.User, error)
	SearchUsers(ctx context.Context, query string, userIDs []uint32) ([]*domain.User, error)
	GetContacts(ctx context.Context, query string, excludeIDs []uint32, currentUserID uint32) ([]*domain.User, error)
}
