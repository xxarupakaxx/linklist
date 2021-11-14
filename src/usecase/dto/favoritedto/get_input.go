package favoritedto

type GetInput struct {
	ReplyToken string `json:"reply_token"`
	LineUserID string `json:"line_user_id"`
}
