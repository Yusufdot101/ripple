package ports

import "context"

type UserVerifier interface {
	VerifyUsers(ctx context.Context, userIDs []uint) (bool, error)
}
