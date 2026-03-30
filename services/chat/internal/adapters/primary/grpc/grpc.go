package grpc

import "context"

func (a *Adapter) VerifyUsers(ctx context.Context, userIDs []uint) (bool, error)
