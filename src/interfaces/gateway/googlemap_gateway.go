package gateway

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/xxarupakaxx/linklist/src/domain/model"
	"googlemaps.github.io/maps"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type GoogleMapGateway struct {
	gmc *maps.Client
}

const (
	maxDetailsOfSearch   = 4
	maxDetailsOfFavorite = 10
	photoAPIURL          = "https://maps.googleapis.com/maps/api/place/photo?maxwidth=400&photoreference="
	noImageURL           = "https://1.bp.blogspot.com/-D2I7Z7-HLGU/Xlyf7OYUi8I/AAAAAAABXq4/jZ0035aDGiE5dP3WiYhlSqhhMgGy8p7zACNcBGAsYHQ/s1600/no_image_square.jpg"
)

func NewGoogleMapGateway() *GoogleMapGateway {
	apiKey := os.Getenv("GMAP_API_KEY")
	gmc, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		logrus.Fatalf("error creating googleMap :%w", err)
	}
	return &GoogleMapGateway{gmc: gmc}
}
func (gateway *GoogleMapGateway) GetPlaceDetailAndPhotoURLsWithQuery(q string) []model.Place {
	places :=gateway.searchPlacesWithQuery(q)
	placeIDs := gateway.getPlaceIDs(places.Results)

	return gateway.GetPlaceDetailsAndPhotoURLs(placeIDs,false)
}

func (gateway *GoogleMapGateway) GetPlaceDetailsAndPhotoURLsWithQueryLatLng(q string, lat, lng float64) []model.Place {
	places := gateway.searchPlacesWithQueryLatLng(q,lat,lng)
	placeIDs := gateway.getPlaceIDs(places.Results)

	return gateway.GetPlaceDetailsAndPhotoURLs(placeIDs,false)
}

func (gateway *GoogleMapGateway) GetPlaceDetailsAndPhotoURLs(placeIDs []string, isFavorite bool) []model.Place {
	googleMapOutputs := []model.Place{}

	maxDetails := maxDetailsOfSearch
	if isFavorite {
		maxDetails = maxDetailsOfFavorite
	}

	for i, placeID := range placeIDs {
		if i == maxDetails {
			break
		}

		placeDetail := gateway.getPlaceDetail(placeID)
		placePhotoURL := noImageURL
		if len(placeDetail.Photos) > 0 {
			placePhotoURL = gateway.getPlacePhotoURL(placeDetail.Photos[0].PhotoReference)
		}

		googleMapOutput := model.Place{
			Name:      placeDetail.Name,
			PlaceID:   placeDetail.PlaceID,
			Address:   placeDetail.FormattedAddress,
			URL:       placeDetail.URL,
			PhotoURL:  placePhotoURL,
		}

		googleMapOutputs = append(googleMapOutputs,googleMapOutput)
	}

	return googleMapOutputs
}

func (gateway *GoogleMapGateway) getPlaceIDs(places []maps.PlacesSearchResult) []string {
	placeIDs := []string{}
	for _, p := range places {
		placeIDs = append(placeIDs, p.PlaceID)
	}
	return placeIDs
}

func (gateway *GoogleMapGateway) searchPlacesWithQuery(q string) maps.PlacesSearchResponse {
	r := &maps.TextSearchRequest{
		Query: q,
		Location: &maps.LatLng{
			Lat: 35.658517,
			Lng: 139.70133399999997,
		},
		Radius:   10,
		Language: "ja",
	}

	res, err := gateway.gmc.TextSearch(context.TODO(), r)
	if err != nil {
		logrus.Errorf("error GoogleMap TextSearch:%v", err)
		res = maps.PlacesSearchResponse{}
	}
	return res
}

func (gateway *GoogleMapGateway) searchPlacesWithQueryLatLng(q string, lat, lng float64) maps.PlacesSearchResponse {
	r := &maps.TextSearchRequest{
		Query:    q,
		Language: "ja",
		Location: &maps.LatLng{Lat: lat, Lng: lng},
		Radius:   500,
	}

	res, err := gateway.gmc.TextSearch(context.Background(), r)
	if err != nil {
		logrus.Errorf("Error GoogleMap TextSearch: %v", err)
		res = maps.PlacesSearchResponse{}
	}
	return res
}

func (gateway *GoogleMapGateway) getPlaceDetail(placeID string) maps.PlaceDetailsResult {
	nameFM, _ := maps.ParsePlaceDetailsFieldMask("name")
	placeIDFM, _ := maps.ParsePlaceDetailsFieldMask("place_id")
	addrFM, _ := maps.ParsePlaceDetailsFieldMask("formatted_address")
	urlFM, _ := maps.ParsePlaceDetailsFieldMask("url")
	photoFM, _ := maps.ParsePlaceDetailsFieldMask("photo")

	r := &maps.PlaceDetailsRequest{
		PlaceID:  placeID,
		Language: "ja",
		Fields: []maps.PlaceDetailsFieldMask{
			nameFM,
			placeIDFM,
			addrFM,
			urlFM,
			photoFM,
		},
	}

	res,err := gateway.gmc.PlaceDetails(context.TODO(),r)
	if err != nil {
		logrus.Errorf("error googleMap placeDetails : %w",err)
		res = maps.PlaceDetailsResult{}
	}
	return res
}

func (gateway *GoogleMapGateway) getPlacePhotoURL(photoReference string) string {
	targetURL := photoAPIURL + photoReference +"&key=" + os.Getenv("GMAP_API_KEY")

	RedirectAttemptedError := errors.New("redirect")
	client :=&http.Client{Timeout: time.Duration(3)*time.Second,CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return RedirectAttemptedError
	}}

	resp,err := client.Head(targetURL)
	if urlError,ok :=err.(*url.Error);ok&& urlError.Err ==RedirectAttemptedError {
		return resp.Header["location"][0]
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Fatalf("failed close body:%w",err)
		}
	}(resp.Body)
	logrus.Errorf("Error GoogleMap getPlacePhotoURL: %v", err)
	return noImageURL
}