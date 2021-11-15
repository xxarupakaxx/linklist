package favoritedto

type AddOutput struct {
	ReplyToken     string `json:"reply_token"`
	UserExists     bool   `json:"user_exists"`
	IsAlreadyAdded bool   `json:"is_already_added"`
}
