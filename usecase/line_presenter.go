package usecase

import (
	"github.com/xxarupakaxx/linklist/usecase/output"
)

type ILinePresenter interface {
	AddFavorite(output output.Add)
	GetFavorites(output output.Get)
	RemoveFavorite(output output.Remove)
	Search(output output.Search)
}
