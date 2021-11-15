package presenter

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
	"github.com/xxarupakaxx/linklist/src/domain/model"
	"github.com/xxarupakaxx/linklist/src/usecase/dto/favoritedto"
	"github.com/xxarupakaxx/linklist/src/usecase/dto/searchdto"
	"os"
	"unicode/utf8"
)

const (
	msgFailAddFavorite       = "お気に入り登録できませんでした"
	msgAlreadyAddAddFavorite = "既に追加済みです！"
	msgSuccssAddFavorite     = "お気に入りに追加しました！"
)

const (
	msgNoRegistGetFavorites            = "お気に入り登録されていません"
	msgAltTextGetFavorites             = "お気に入り一覧の表示結果です"
	msgPostbackActionLabelGetFavorites = "Remove"
	msgPostbackActionDataGetFavorites  = "action=removeFavorite&placeId=%s"
)

const (
	msgFailRemoveFavorite       = "お気に入りを削除できませんでした。"
	msgAlreadyAddRemoveFavorite = "すでに削除済みです"
	msgSuccessRemoveFavorite    = "お気に入りを削除しました"
)

const (
	msgNoRegisterSearch          = "検索結果は0件でした"
	msgAltTextSearch             = "「%s」の検索結果です"
	msgPostbackActionLabelSearch = "add to my favorites"
	msgPostbackActionDataSearch  = "action=addFavorite&placeId=%s"
)

const maxTextWC = 60

type LinePresenter struct {
	bot *linebot.Client
}

type carouselMsgs struct {
	noResult            string
	altText             string
	postbackActionLabel string
	postbackActionData  string
}

func NewLinePresenter() *LinePresenter {
	secret := os.Getenv("LINEBOT_SECRET")
	token := os.Getenv("LINEBOT_TOKEN")

	bot, err := linebot.New(secret, token)
	if err != nil {
		logrus.Fatalf("Error creating LINEBOT client:%v", err)
	}
	return &LinePresenter{bot: bot}
}

func (l *LinePresenter) addFavorite(output favoritedto.AddOutput) {
	replyToken := output.ReplyToken

	if !output.UserExists {
		l.replyMessage(replyToken,msgFailAddFavorite)
	} else if output.IsAlreadyAdded {
		l.replyMessage(replyToken,msgAlreadyAddAddFavorite)
	} else {
		l.replyMessage(replyToken, msgSuccssAddFavorite)
	}
}

func (l *LinePresenter) GetFavorites(output favoritedto.GetOutput) {
	msgs :=carouselMsgs{
		noResult:            msgNoRegistGetFavorites,
		altText:             msgAltTextGetFavorites,
		postbackActionLabel: msgPostbackActionLabelGetFavorites,
		postbackActionData:  msgPostbackActionDataGetFavorites,
	}
	l.replyCarouselColumn(msgs,output.GoogleMapOutputs,output.ReplyToken)
}

func (l *LinePresenter) RemoveFavorite(output favoritedto.RemoveOutput) {
	if !output.UserExists {
		l.replyMessage(output.ReplyToken,msgFailRemoveFavorite)
	}else if output.IsAlreadyRemoved {
		l.replyMessage(output.ReplyToken,msgAlreadyAddRemoveFavorite)
	}else {
		l.replyMessage(msgSuccessRemoveFavorite,output.ReplyToken)
	}
}

func (l *LinePresenter) Search(output searchdto.Output) {
	msgs := carouselMsgs{
		noResult:            msgNoRegisterSearch,
		altText:             fmt.Sprintf(msgAltTextSearch, output.Q),
		postbackActionLabel: msgPostbackActionLabelSearch,
		postbackActionData:  msgPostbackActionDataSearch,
	}
	l.replyCarouselColumn(msgs,output.GoogleMapOutputs,output.ReplyToken)
}

func (l *LinePresenter) replyMessage(replyToken, msg string) {
	res := linebot.NewTextMessage(msg)
	if _,err := l.bot.ReplyMessage(replyToken,res).Do();err != nil {
		logrus.Errorf("error replying message : %w",err)
	}
}

func (l *LinePresenter) replyCarouselColumn(msg carouselMsgs, googleMapOutputs []model.Place, replyToken string) {
	if len(googleMapOutputs) == 0 {
		l.replyMessage(replyToken,msg.noResult)
		return
	}

	ccs :=[]*linebot.CarouselColumn{}
	for _, output := range googleMapOutputs {
		addr := output.Address
		if maxTextWC < utf8.RuneCountInString(addr){
			addr = string([]rune(addr)[:maxTextWC])
		}

		data := fmt.Sprintf(msg.postbackActionData,output.PlaceID)
		cc := linebot.NewCarouselColumn(output.PhotoURL,output.Name,addr,linebot.NewURIAction("Open Google Map",output.URL),linebot.NewPostbackAction(msg.postbackActionLabel,data, "","")).WithImageOptions("#FFFFFF")
		ccs = append(ccs,cc)
	}

	res := linebot.NewTemplateMessage(msg.altText,linebot.NewCarouselTemplate(ccs...).WithImageOptions(linebot.ImageAspectRatioTypeRectangle,linebot.ImageSizeTypeCover))
	if _,err:= l.bot.ReplyMessage(replyToken,res).Do();err != nil{
		logrus.Errorf("Error linebot replying message: %w",err)
	}
}