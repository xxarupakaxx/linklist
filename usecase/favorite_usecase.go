package usecase

import (
	"github.com/xxarupakaxx/linklist/usecase/input"
	"github.com/xxarupakaxx/linklist/usecase/output"
)

type IFavoriteUseCase interface {
	Get(input input.Get) output.Get
	Add(input input.Add) output.Add
	Remove(input input.Remove) output.Remove
}
