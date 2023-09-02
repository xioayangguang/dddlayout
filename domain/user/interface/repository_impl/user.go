package repository_impl

import (
	"context"
)

// 当前领域关心的模型
type User struct {
	Id             uint64
	CreatedAt      uint
	UpdatedAt      uint
	DeletedAt      uint
	Uuid           uint
	Serial         uint
	Nickname       string
	Mail           string
	Describe       string
	Code           string
	InvitationCode string
	Avatar         string
	Status         int
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uint64) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetMaxSerial(ctx context.Context) int
}
