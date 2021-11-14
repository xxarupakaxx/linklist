package presenter

import (
	"github.com/xxarupakaxx/linklist/src/usecase/dto/favoritedto"
	"github.com/xxarupakaxx/linklist/src/usecase/dto/searchdto"
)

type ILinePresenter interface {
	AddFavorite(output favoritedto.AddOutput)
	GetFavorites(output favoritedto.GetOutput)
	RemoveFavorite(output favoritedto.RemoveOutput)
	Search(output searchdto.Output)
}
