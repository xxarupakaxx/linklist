package database

import (
	"context"
	"errors"
	"github.com/xxarupakaxx/linklist/domain/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Save(ctx context.Context, lineUserID string) uint {
	tx := ur.db.WithContext(ctx)
	user := model.User{}
	err := tx.Table("users").Where(model.User{LineUserID: lineUserID}).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = model.User{LineUserID: lineUserID}
		tx.Create(&user)
	}
	return user.ID
}

func (ur *UserRepository) FindOne(ctx context.Context, lineUserID string) uint {
	tx := ur.db.WithContext(ctx)
	user := model.User{}
	err := tx.Table("users").Where(model.User{LineUserID: lineUserID}).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0
	}
	return user.ID
}
