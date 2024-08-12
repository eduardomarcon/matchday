package port

import (
	"api/internal/core/domain"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	ExistsUserByEmail(ctx context.Context, email string) bool
}

type UserService interface {
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
}
