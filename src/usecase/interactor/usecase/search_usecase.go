package usecase

import "github.com/xxarupakaxx/linklist/src/usecase/dto/searchdto"

type ISearchUseCase interface {
	Hundle(input searchdto.Input) searchdto.Output
}