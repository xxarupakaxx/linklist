package interactor

import (
	"github.com/sirupsen/logrus"
	 "github.com/xxarupakaxx/linklist/domain/model"
	 "github.com/xxarupakaxx/linklist/usecase/dto/searchdto"
	"github.com/xxarupakaxx/linklist/usecase/gateway"
	"github.com/xxarupakaxx/linklist/usecase/presenter"
	"os"
)

type SearchInteract struct {
	googleMapGateway gateway.IGoogleMapGateway
	linePresenter    presenter.ILinePresenter
}

func (si *SearchInteract) Hundle(input searchdto.Input) searchdto.Output {
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

	output := searchdto.Output{
		ReplyToken:       input.ReplyToken,
		Q:                outQ,
		GoogleMapOutputs: googleMapOutputs,
	}
	if output.ReplyToken != "" {
		si.linePresenter.Search(output)
	}
	return output
}

func NewSearchInteract(googleMapGateway gateway.IGoogleMapGateway, linePresenter presenter.ILinePresenter) *SearchInteract {
	return &SearchInteract{googleMapGateway: googleMapGateway, linePresenter: linePresenter}
}

func isNomination(q string, lat float64, lng float64) bool {
	return q != "" && lat == 0 && lng == 0
}

func isOnlyLocaleInfo(addr string, lat float64, lng float64) bool {
	return addr != "" && lat != 0 && lng != 0
}
