package usecase

import (
	searchdto2 "github.com/xxarupakaxx/linklist/usecase/dto/searchdto"
)

type ISearchUseCase interface {
	Hundle(input searchdto2.Input) searchdto2.Output
}