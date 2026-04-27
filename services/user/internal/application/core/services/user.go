package services

import (
	"context"
	"regexp"

	"github.com/Yusufdot101/ripple/services/user/internal/application/core/domain"
	"github.com/Yusufdot101/ripple/services/user/internal/ports"
)

var EmailRX = regexp.MustCompile(
	"^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
)

func isValidEmail(email string) bool {
	return EmailRX.MatchString(email)
}

type UserService struct {
	repo ports.Repository
}

func NewUserService(repo ports.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (usvc *UserService) GetUsersByEmail(email string) ([]*domain.User, error) {
	return usvc.repo.FindUsersByEmail(email)
}

func (usvc *UserService) GetUsersByIDs(ctx context.Context, userIDs []uint32) ([]*domain.User, error) {
	return usvc.repo.FindUsersByID(ctx, userIDs)
}

func (usvc *UserService) SearchUsers(ctx context.Context, query string, userIDs []uint32) ([]*domain.User, error) {
	return usvc.repo.SearchUsers(ctx, query, userIDs)
}

func (usvc *UserService) GetContacts(ctx context.Context, query string, excludeIDs []uint32, currentUserID uint32) ([]*domain.User, error) {
	return usvc.repo.GetContacts(ctx, query, excludeIDs, currentUserID)
}
