package jsons

type TickerJson struct {
	UserId   string `json:"user_id"`
	Symbol   string `json:"symbol"`
	FcmToken string `json:"fcm_token"`
}
