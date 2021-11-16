package input

type Search struct {
	ReplyToken string  `json:"reply_token"`
	Q          string  `json:"q"`
	Addr       string  `json:"addr"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
}
