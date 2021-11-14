package repository

import "context"

type IFavoriteRepository interface {
	FindAll(ctx context.Context, lineUserID string) []string
	Save(ctx context.Context, id uint, placeID string) bool
	Delete(ctx context.Context, id uint, placeID string) bool
}
