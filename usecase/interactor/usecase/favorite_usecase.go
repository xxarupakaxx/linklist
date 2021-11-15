package usecase

import (
	 "github.com/xxarupakaxx/linklist/usecase/dto/favoritedto"
)

type IFavoriteUseCase interface {
	Get(input favoritedto.GetInput) favoritedto.GetOutput
	Add(input favoritedto.AddInput) favoritedto.AddOutput
	Remove(input favoritedto.RemoveInput) favoritedto.RemoveOutput
}
