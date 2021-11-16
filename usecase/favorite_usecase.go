package usecase

import (
	"context"
	"github.com/xxarupakaxx/linklist/domain/repository"
	"github.com/xxarupakaxx/linklist/usecase/input"
	"github.com/xxarupakaxx/linklist/usecase/output"
)

type IFavoriteUseCase interface {
	Get(input input.Get) output.Get
	Add(input input.Add) output.Add
	Remove(input input.Remove) output.Remove
}

type FavoriteInteracter struct {
	userRepository     repository.IUserRepository
	favoriteRepository repository.IFavoriteRepository
	googleMapGateway   IGoogleMapGateway
	linePresenter      ILinePresenter
	context            context.Context
}

func (interact *FavoriteInteracter) Get(input input.Get) output.Get {
	placeIDs := interact.favoriteRepository.FindAll(interact.context, input.LineUserID)
	googleMapOutputs := interact.googleMapGateway.GetPlaceDetailsAndPhotoURLs(placeIDs, true)

	output := output.Get{
		ReplyToken:       input.ReplyToken,
		GoogleMapOutputs: googleMapOutputs,
	}
	if output.ReplyToken != "" {
		interact.linePresenter.GetFavorites(output)
	}

	return output
}

func (interact *FavoriteInteracter) Add(input input.Add) output.Add {
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

	output := output.Add{
		ReplyToken:     input.ReplyToken,
		UserExists:     userExists,
		IsAlreadyAdded: !isAdded,
	}

	if output.ReplyToken != "" {
		interact.linePresenter.AddFavorite(output)
	}

	return output
}

func (interact *FavoriteInteracter) Remove(input input.Remove) output.Remove {
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

	output := output.Remove{
		ReplyToken:       input.ReplyToken,
		UserExists:       userExists,
		IsAlreadyRemoved: !isRemoved,
	}
	if output.ReplyToken != "" {
		interact.linePresenter.RemoveFavorite(output)
	}

	return output
}

func NewFavoriteInteract(userRepository repository.IUserRepository, favoriteRepository repository.IFavoriteRepository, googleMapGateway IGoogleMapGateway, linePresenter ILinePresenter) *FavoriteInteracter {
	return &FavoriteInteracter{
		userRepository:     userRepository,
		favoriteRepository: favoriteRepository,
		googleMapGateway:   googleMapGateway,
		linePresenter:      linePresenter,
	}
}