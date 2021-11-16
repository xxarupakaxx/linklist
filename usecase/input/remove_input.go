package input

type Remove struct {
	ReplyToken string `json:"reply_token"`
	LineUserID string `json:"line_user_id"`
	PlaceID    string `json:"place_id"`
}
