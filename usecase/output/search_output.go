package output

import (
	"github.com/xxarupakaxx/linklist/domain/model"
)

type Search struct {
	ReplyToken       string         `json:"reply_token"`
	Q                string         `json:"q"`
	GoogleMapOutputs []model.Place `json:"google_map_outputs"`
}
