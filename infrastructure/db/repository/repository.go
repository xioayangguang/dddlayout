package repository

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
	"layout/infrastructure/db/connect"
)

var ProviderSet = wire.NewSet(
	connect.NewDB,
	NewRepository,
	NewUserRepository,
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) getDb(ctx context.Context) (dbHandler *gorm.DB) {
	dbHandler, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		dbHandler = r.db
	}
	return dbHandler.WithContext(ctx)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
