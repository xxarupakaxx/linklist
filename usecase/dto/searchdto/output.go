package searchdto

import (
	"github.com/xxarupakaxx/linklist/domain/model"
)

type Output struct {
	ReplyToken       string         `json:"reply_token"`
	Q                string         `json:"q"`
	GoogleMapOutputs []model.Place `json:"google_map_outputs"`
}
