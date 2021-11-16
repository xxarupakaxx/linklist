package output

import "github.com/xxarupakaxx/linklist/domain/model"

type Add struct {
	ReplyToken     string `json:"reply_token"`
	UserExists     bool   `json:"user_exists"`
	IsAlreadyAdded bool   `json:"is_already_added"`
}

type Get struct {
	ReplyToken       string        `json:"reply_token"`
	GoogleMapOutputs []model.Place `json:"google_map_outputs"`
}

type Remove struct {
	ReplyToken       string `json:"reply_token"`
	UserExists       bool   `json:"user_exists"`
	IsAlreadyRemoved bool   `json:"is_already_removed"`
}
