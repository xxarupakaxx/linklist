package gateway

import "github.com/xxarupakaxx/linklist/src/domain/model"

type IGoogleMapGateway interface {
	GetPlaceDetailAndPhotoURLsWithQuery(string2 string) []model.Place
	GetPlaceDetailsAndPhotoURLsWithQueryLatLng(string2 string,lat,lng float64) []model.Place
	GetPlaceDetailsAndPhotoURLs([]string,bool) []model.Place
}
