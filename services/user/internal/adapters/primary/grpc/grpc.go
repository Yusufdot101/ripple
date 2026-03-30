package grpc

import (
	"context"

	"github.com/Yusufdot101/ripple-proto/golang/user"
)

func (a *Adapter) VerifyUsers(context.Context, *user.VerifyUsersRequest) (*user.VerifyUsersResponse, error) {
	return &user.VerifyUsersResponse{
		AllValid: true,
	}, nil
}
