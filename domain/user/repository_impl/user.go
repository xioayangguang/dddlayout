package repository_impl

import (
	"context"
	"layout/domain/user/model/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	Update(ctx context.Context, user *entities.User) error
	GetByID(ctx context.Context, id uint64) (*entities.User, error)
	GetByUsername(ctx context.Context, username string) (*entities.User, error)
	GetMaxSerial(ctx context.Context) int
}
