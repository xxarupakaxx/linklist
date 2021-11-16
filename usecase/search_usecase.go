package usecase

import (
	"github.com/xxarupakaxx/linklist/usecase/input"
	"github.com/xxarupakaxx/linklist/usecase/output"
)

type ISearchUseCase interface {
	Hundle(input input.Search) output.Search
}