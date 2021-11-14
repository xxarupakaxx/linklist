package favoritedto

import "github.com/xxarupakaxx/linklist/src/domain/model"

type GetOutput struct {
	ReplyToken       string        `json:"reply_token"`
	GoogleMapOutputs []model.Place `json:"google_map_outputs"`
}
