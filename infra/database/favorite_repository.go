package database

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/xxarupakaxx/linklist/domain/model"
	"gorm.io/gorm"
)

type FavoriteRepository struct {
	db *gorm.DB
}

func (fr *FavoriteRepository) FindAll(ctx context.Context, lineUserID string) []string {
	tx := fr.db.WithContext(ctx)
	user := model.User{}
	tx.Table("users").Where(model.User{LineUserID: lineUserID}).First(&user)

	favorites := []model.Favorite{}
	if err := tx.Model(&user).Select("favorite").Scan(&favorites).Error; err != nil {
		logrus.Fatalf("Favoriteスキャンできない ", err)
	}

	placeIDs := []string{}

	for _, f := range favorites {
		placeIDs = append(placeIDs, f.PlaceID)
	}

	return placeIDs

}

func (fr *FavoriteRepository) Save(ctx context.Context, id uint, placeID string) bool {
	tx := fr.db.WithContext(ctx)
	favorite := model.Favorite{}
	err := tx.Table("favorites").Where(model.Favorite{UserID: id, PlaceID: placeID}).First(&favorite).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		favorite = model.Favorite{UserID: id, PlaceID: placeID}
		tx.Create(&favorite)
		return true
	}
	return false
}

func (fr *FavoriteRepository) Delete(ctx context.Context, id uint, placeID string) bool {
	tx := fr.db.WithContext(ctx)
	favorite := model.Favorite{}
	err := tx.Table("favorite").Where(model.Favorite{UserID: id, PlaceID: placeID}).First(&favorite).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	tx.Delete(&favorite)
	return true
}

func NewFavoriteRepository(db *gorm.DB) *FavoriteRepository {
	return &FavoriteRepository{db: db}
}
