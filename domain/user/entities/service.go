package entities

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
	"layout/domain/user/service"
)

var ProviderSet = wire.NewSet(
	NewService,
	service.NewUserService,
)

type Service struct {
	db *gorm.DB
}

// todo  事务不应该在此开启，应该在application 层
func (s *Service) transaction(ctx context.Context, callBack func(ctx context.Context) error) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, "tx", tx)
		return callBack(ctx)
	})
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}
