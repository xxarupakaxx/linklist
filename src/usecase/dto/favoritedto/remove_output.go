package favoritedto

type RemoveOutput struct {
	ReplyToken       string `json:"reply_token"`
	UserExists       bool   `json:"user_exists"`
	IsAlreadyRemoved bool   `json:"is_already_removed"`
}
