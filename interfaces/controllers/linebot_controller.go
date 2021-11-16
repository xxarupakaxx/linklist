package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
	usecase2 "github.com/xxarupakaxx/linklist/usecase"
	"github.com/xxarupakaxx/linklist/usecase/input"
	"net/http"
	"os"
	"strings"
)

type LinebotController struct {
	favoriteInteractor usecase2.IFavoriteUseCase
	searchInteractor   usecase2.SearchInteract
	bot                *linebot.Client
}

func NewLinebotController(favoriteInteractor usecase2.IFavoriteUseCase, searchInteractor usecase2.SearchInteract, bot *linebot.Client) *LinebotController {
	secret := os.Getenv("LINEBOT_SECRET")
	token := os.Getenv("LINEBOT_TOKEN")

	bot,err := linebot.New(secret,token)
	if err != nil {
		logrus.Fatalf("error creating bot:%w",err)
	}
	return &LinebotController{favoriteInteractor: favoriteInteractor, searchInteractor: searchInteractor, bot: bot}
}

func (controller *LinebotController) CatchEvents() echo.HandlerFunc {
	return func(c echo.Context) error {
		events,err := controller.bot.ParseRequest(c.Request())
		if err != nil {
			logrus.Fatalf("error LINEBOT parsing request : %w",err)
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch event.Message.(type) {
				case *linebot.TextMessage:
					controller.replyToTextMessage(event)
				case *linebot.LocationMessage:
					controller.replyToLocationMessage(event)
				}
			}else if event.Type == linebot.EventTypePostback {
				controller.replyToEventTypePostBack(event)
			}
		}
		return c.NoContent(http.StatusOK)
	}
}
func (controller *LinebotController) replyToTextMessage(e *linebot.Event) {
	msg := e.Message.(*linebot.TextMessage).Text

	if msg == "お気に入り" {
		favoriteGetInput := input.Get{
			ReplyToken: e.ReplyToken,
			LineUserID: e.Source.UserID,
		}
		controller.favoriteInteractor.Get(favoriteGetInput)
	}else{
		searchInput := input.Search{
			ReplyToken: e.ReplyToken,
			Q:          msg,
		}
		controller.searchInteractor.Hundle(searchInput)
	}
}

func (controller *LinebotController) replyToLocationMessage(e *linebot.Event) {
	msg := e.Message.(*linebot.LocationMessage)

	searchInput := input.Search{
		ReplyToken: e.ReplyToken,
		Q:          msg.Title,
		Addr:       excerptAddr(msg.Address),
		Lat:        msg.Latitude,
		Lng:        msg.Longitude,
	}
	controller.searchInteractor.Hundle(searchInput)
}

func (controller *LinebotController) replyToEventTypePostBack(e *linebot.Event)  {
	dataMap := createDataMap(e.Postback.Data)
	if dataMap["action"]== "addFavorite" {
		favoriteAddInput := input.Add{
			ReplyToken: e.ReplyToken,
			LineUserID: e.Source.UserID,
			PlaceID:    dataMap["placeId"],
		}
		controller.favoriteInteractor.Add(favoriteAddInput)
	}else if dataMap["action"] == "removeFavorite" {
		favoriteRemoveInput := input.Remove{
			ReplyToken: e.ReplyToken,
			LineUserID: e.Source.UserID,
			PlaceID:    dataMap["placeId"],
		}
		controller.favoriteInteractor.Remove(favoriteRemoveInput)
	}
}

func excerptAddr(fullAddr string) string {
	addrAr :=strings.Split(fullAddr," ")
	return addrAr[1]
}

func createDataMap(q string) map[string]string {
	dataMap :=make(map[string]string)

	dataArr := strings.Split(q,"&")
	for _, s := range dataArr {
		splitedData :=strings.Split(s,"=")
		dataMap[splitedData[0]]=splitedData[1]
	}

	return dataMap
}