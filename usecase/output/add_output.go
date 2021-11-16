package output

type Add struct {
	ReplyToken     string `json:"reply_token"`
	UserExists     bool   `json:"user_exists"`
	IsAlreadyAdded bool   `json:"is_already_added"`
}
