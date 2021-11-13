package repository

type IFavoriteRepository interface {
	FindAll(placeID string) []string
	Save(id uint, placeID string) bool
	Delete(id uint, placeID string) bool
}
