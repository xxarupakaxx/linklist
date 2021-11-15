package interactor

import (
	"context"
	 "github.com/xxarupakaxx/linklist/domain/repository"
	 "github.com/xxarupakaxx/linklist/usecase/dto/favoritedto"
	 "github.com/xxarupakaxx/linklist/usecase/gateway"
	 "github.com/xxarupakaxx/linklist/usecase/presenter"
)

type FavoriteInteract struct {
	userRepository     repository.IUserRepository
	favoriteRepository repository.IFavoriteRepository
	googleMapGateway   gateway.IGoogleMapGateway
	linePresenter      presenter.ILinePresenter
	context            context.Context
}

func (interact *FavoriteInteract) Get(input favoritedto.GetInput) favoritedto.GetOutput {
	placeIDs := interact.favoriteRepository.FindAll(interact.context, input.LineUserID)
	googleMapOutputs := interact.googleMapGateway.GetPlaceDetailsAndPhotoURLs(placeIDs, true)

	output := favoritedto.GetOutput{
		ReplyToken:       input.ReplyToken,
		GoogleMapOutputs: googleMapOutputs,
	}
	if output.ReplyToken != "" {
		interact.linePresenter.GetFavorites(output)
	}

	return output
}

func (interact *FavoriteInteract) Add(input favoritedto.AddInput) favoritedto.AddOutput {
	userID := interact.userRepository.Save(interact.context, input.LineUserID)
	var userExists bool
	var isAdded bool
	if userID == 0 {
		userExists = false
		isAdded = false
	} else {
		userExists = true
		isAdded = interact.favoriteRepository.Save(interact.context, userID, input.PlaceID)
	}

	output := favoritedto.AddOutput{
		ReplyToken:     input.ReplyToken,
		UserExists:     userExists,
		IsAlreadyAdded: !isAdded,
	}

	if output.ReplyToken != "" {
		interact.linePresenter.AddFavorite(output)
	}

	return output
}

func (interact *FavoriteInteract) Remove(input favoritedto.RemoveInput) favoritedto.RemoveOutput {
	userID := interact.userRepository.FindOne(interact.context, input.LineUserID)

	var userExists bool
	var isRemoved bool
	if userID == 0 {
		userExists = false
		isRemoved = false
	} else {
		userExists = true
		isRemoved = interact.favoriteRepository.Delete(interact.context, userID, input.PlaceID)
	}

	output := favoritedto.RemoveOutput{
		ReplyToken:       input.ReplyToken,
		UserExists:       userExists,
		IsAlreadyRemoved: !isRemoved,
	}
	if output.ReplyToken != "" {
		interact.linePresenter.RemoveFavorite(output)
	}

	return output
}

func NewFavoriteInteract(userRepository repository.IUserRepository, favoriteRepository repository.IFavoriteRepository, googleMapGateway gateway.IGoogleMapGateway, linePresenter presenter.ILinePresenter) *FavoriteInteract {
	return &FavoriteInteract{
		userRepository:     userRepository,
		favoriteRepository: favoriteRepository,
		googleMapGateway:   googleMapGateway,
		linePresenter:      linePresenter,
	}
}
