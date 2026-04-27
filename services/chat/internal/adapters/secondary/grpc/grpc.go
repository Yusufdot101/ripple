package grpc

import (
	"context"

	userpb "github.com/Yusufdot101/ripple-proto/golang/user/v4"
)

func (a *Adapter) VerifyUsers(ctx context.Context, userIDs []uint) (bool, error) {
	userIDs32 := []uint32{}
	for _, userID := range userIDs {
		userIDs32 = append(userIDs32, uint32(userID))
	}

	req := &userpb.VerifyUsersRequest{
		Ids: userIDs32,
	}
	res, err := a.userClient.VerifyUsers(ctx, req)
	if err != nil {
		return false, err
	}
	return res.AllValid, nil
}

func (a *Adapter) GetUsersByIDs(ctx context.Context, userIDs []uint) ([]*userpb.User, error) {
	userIDs32 := []uint32{}
	for _, userID := range userIDs {
		userIDs32 = append(userIDs32, uint32(userID))
	}

	req := &userpb.GetUsersRequest{
		Ids: userIDs32,
	}
	res, err := a.userClient.GetUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Users, nil
}

func (a *Adapter) SearchUsers(ctx context.Context, query string, userIDs []uint) ([]*userpb.User, error) {
	userIDs32 := []uint32{}
	for _, userID := range userIDs {
		userIDs32 = append(userIDs32, uint32(userID))
	}

	req := &userpb.SearchUsersRequest{
		Ids:   userIDs32,
		Query: query,
	}
	res, err := a.userClient.SearchUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Users, nil
}

func (a *Adapter) GetContacts(ctx context.Context, query string, excludeIDs []uint, currentUserID uint) ([]*userpb.User, error) {
	excludeIDsUint := []uint32{}
	for _, id := range excludeIDs {
		excludeIDsUint = append(excludeIDsUint, uint32(id))
	}
	req := &userpb.GetContactsRequest{
		UserId:     uint32(currentUserID),
		ExcludeIds: excludeIDsUint,
		Query:      query,
	}
	res, err := a.userClient.GetContacts(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Users, nil
}
