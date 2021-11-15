package controllers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
	"github.com/xxarupakaxx/linklist/usecase/dto/favoritedto"
	"github.com/xxarupakaxx/linklist/usecase/dto/searchdto"
	"github.com/xxarupakaxx/linklist/usecase/interactor/usecase"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const msgSetPram = "パラメータを設定指定ください"

type APIController struct {
	favoriteInteractor usecase.IFavoriteUseCase
	searchInteractor   usecase.ISearchUseCase
	bot                *linebot.Client
}

func NewAPIController(favoriteInteractor usecase.IFavoriteUseCase, searchInteractor usecase.ISearchUseCase) *APIController {
	return &APIController{favoriteInteractor: favoriteInteractor, searchInteractor: searchInteractor}
}

func (controller *APIController) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		q := c.QueryParam("q")
		addr := c.QueryParam("addr")
		latStr := c.QueryParam("lat")
		lngStr := c.QueryParam("lng")

		if q == "" && (addr == "" || latStr == "" || lngStr == "") {
			return c.JSON(http.StatusBadRequest, msgSetPram)
		}

		lat, lng := float64(0), float64(0)
		if latStr != "" && lngStr != "" {
			var err error
			lat, err = strconv.ParseFloat(latStr, 64)
			lng, err = strconv.ParseFloat(lngStr, 64)
			if err != nil {
				logrus.Error("error parse float:%w", err)
			}
		}

		input := searchdto.Input{
			Q:    q,
			Addr: addr,
			Lat:  lat,
			Lng:  lng,
		}
		output := controller.searchInteractor.Hundle(input)

		return c.JSON(http.StatusOK, output)
	}
}

func (controller *APIController) GetFavorites() echo.HandlerFunc {
	return func(c echo.Context) error {
		lineIDToken := c.FormValue("line_id_token")
		lineUserID := getLineUserIDByToken(lineIDToken)
		if lineUserID == "" {
			return c.JSON(http.StatusBadRequest, msgSetPram)
		}

		input := favoritedto.GetInput{LineUserID: lineUserID}
		output := controller.favoriteInteractor.Get(input)
		return c.JSON(http.StatusOK, output)
	}
}

func (controller *APIController) AddFavorites() echo.HandlerFunc {
	return func(c echo.Context) error {
		lineIDToken := c.FormValue("line_id_token")
		lineUseID := getLineUserIDByToken(lineIDToken)
		placeID := c.FormValue("place_id")

		if lineUseID == "" || placeID == "" {
			return c.JSON(http.StatusBadRequest, msgSetPram)
		}

		input := favoritedto.AddInput{
			LineUserID: lineUseID,
			PlaceID:    placeID,
		}
		output :=controller.favoriteInteractor.Add(input)

		return c.JSON(http.StatusOK,output)
	}
}

func (controller *APIController) RemoveFavorites() echo.HandlerFunc {
	return func(c echo.Context) error {
		lineIDToken := c.FormValue("line_id_token")

		lineUserID := getLineUserIDByToken(lineIDToken)
		placeID := c.FormValue("place_id")

		if lineUserID ==""||placeID ==""{
			return c.JSON(http.StatusBadRequest, msgSetPram)
		}

		input := favoritedto.RemoveInput{
			LineUserID: lineUserID,
			PlaceID:    placeID,
		}
		output :=controller.favoriteInteractor.Remove(input)

		return c.JSON(http.StatusOK,output)
	}
}

type verifyResp struct {
	Sub string `json:"sub"`
}

func getLineUserIDByToken(idToken string) string {
	values := url.Values{}
	values.Add("id_token", idToken)
	values.Add("client_id", os.Getenv("LIFF_CHANNEL_ID"))

	resp, err := http.PostForm("https://api.line.me/oauth2/v2.1/verify", values)

	if err != nil {
		logrus.Errorf("error Parsing LINEIDToken :%w", err)
		return ""
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Println(err)
		}
	}(resp.Body)
	jsonBytes := ([]byte)(string(body))
	data := new(verifyResp)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		logrus.Errorf("error JSON Unmarshal:%w", err)
		return ""
	}
	return data.Sub
}
