package repository

import (
	"context"
	"github.com/pkg/errors"
	"layout/domain/user/model/entities"
	"layout/infrastructure/db/model"
)

type userRepository struct {
	*Repository
}

func NewUserRepository(r *Repository) facade.UserRepository {
	return &userRepository{
		Repository: r,
	}
}
func (r *userRepository) Create(ctx context.Context, user *entities.User) error {
	if err := r.getDb(ctx).Create(user).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *entities.User) error {
	if err := r.getDb(ctx).Save(user).Error; err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, userId uint64) (*entities.User, error) {
	var user model.User
	var user1 entities.User
	if err := r.getDb(ctx).Model(user).Where("id = ?", userId).First(&user1).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get user by ID")
	}
	return &user1, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user model.User
	var user1 entities.User
	err := r.getDb(ctx).Model(user).Where("nickname = ?", username).First(&user1).Error
	return &user1, err
}

func (r *userRepository) GetMaxSerial(ctx context.Context) int {
	var u model.User
	r.getDb(ctx).Order("serial desc").Take(&u)
	return int(u.Serial)
}
