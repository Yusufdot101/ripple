package grpc

import (
	"context"
	"errors"

	"github.com/Yusufdot101/ripple-proto/golang/user"
	"github.com/Yusufdot101/ripple/services/user/internal/application/core/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Adapter) VerifyUsers(ctx context.Context, req *user.VerifyUsersRequest) (*user.VerifyUsersResponse, error) {
	allValid, err := a.asvc.VerifyUsers(ctx, req.Ids)
	if err != nil {
		var st error
		if errors.Is(err, domain.ErrInvalidID) {
			st = status.Error(codes.NotFound, err.Error())
		} else {
			st = status.Error(codes.Internal, err.Error())
		}
		return &user.VerifyUsersResponse{
			AllValid: false,
		}, st
	}

	return &user.VerifyUsersResponse{
		AllValid: allValid,
	}, nil
}
