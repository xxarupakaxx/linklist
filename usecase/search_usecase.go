package usecase

import (
	"github.com/sirupsen/logrus"
	"github.com/xxarupakaxx/linklist/domain/model"
	"github.com/xxarupakaxx/linklist/usecase/input"
	"github.com/xxarupakaxx/linklist/usecase/output"
	"os"
)

type ISearchUseCase interface {
	Hundle(input input.Search) output.Search
}

type SearchInteractor struct {
	googleMapGateway IGoogleMapGateway
	linePresenter    ILinePresenter
}

func (si *SearchInteractor) Hundle(input input.Search) output.Search {
	outQ := ""
	var googleMapOutputs []model.Place
	if isNomination(input.Q, input.Lat, input.Lng) {
		outQ = input.Q
		q := outQ + " " + os.Getenv("QUERY")
		googleMapOutputs = si.googleMapGateway.GetPlaceDetailAndPhotoURLsWithQuery(q)
	} else if isOnlyLocaleInfo(input.Addr, input.Lat, input.Lng) {
		outQ = input.Addr
		q := os.Getenv("QUERY") + " " + outQ
		googleMapOutputs = si.googleMapGateway.GetPlaceDetailsAndPhotoURLsWithQueryLatLng(q, input.Lat, input.Lng)
	} else {
		logrus.Error("Error　unexpected user request")
	}

	search := output.Search{
		ReplyToken:       input.ReplyToken,
		Q:                outQ,
		GoogleMapOutputs: googleMapOutputs,
	}
	if search.ReplyToken != "" {
		si.linePresenter.Search(search)
	}
	return search
}

func NewSearchInteract(googleMapGateway IGoogleMapGateway, linePresenter ILinePresenter) *SearchInteractor {
	return &SearchInteractor{googleMapGateway: googleMapGateway, linePresenter: linePresenter}
}

func isNomination(q string, lat float64, lng float64) bool {
	return q != "" && lat == 0 && lng == 0
}

func isOnlyLocaleInfo(addr string, lat float64, lng float64) bool {
	return addr != "" && lat != 0 && lng != 0
}
