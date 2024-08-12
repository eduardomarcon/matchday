package service

import (
	"api/internal/core/domain"
	"api/internal/core/port"
	"context"
	"log/slog"
)

type UserService struct {
	repository port.UserRepository
}

func NewUserService(repository port.UserRepository) *UserService {
	return &UserService{
		repository,
	}
}

func (us *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	existingUser := us.repository.ExistsUserByEmail(ctx, user.Email)
	if existingUser {
		slog.Error("email is already linked to a user", "email", user.Email)
		return nil, domain.ErrConflictingData
	}

	insertedUser, err := us.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return insertedUser, nil
}
