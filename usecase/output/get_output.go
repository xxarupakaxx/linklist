package output

import (
	"github.com/xxarupakaxx/linklist/domain/model"
)

type Get struct {
	ReplyToken       string         `json:"reply_token"`
	GoogleMapOutputs []model.Place `json:"google_map_outputs"`
}
