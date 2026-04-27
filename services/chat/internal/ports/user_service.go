package ports

import (
	"context"

	userpb "github.com/Yusufdot101/ripple-proto/golang/user/v4"
)

type UserService interface {
	VerifyUsers(ctx context.Context, userIDs []uint) (bool, error)
	GetUsersByIDs(ctx context.Context, userIDs []uint) ([]*userpb.User, error)
	SearchUsers(ctx context.Context, query string, userIDs []uint) ([]*userpb.User, error)
	GetContacts(ctx context.Context, query string, excludeIDs []uint, currentUserID uint) ([]*userpb.User, error)
}
