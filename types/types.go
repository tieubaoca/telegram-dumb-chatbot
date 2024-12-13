package types

type Message struct {
	ChatId  int64  `json:"chat_id"`
	From    int64  `json:"from"`
	Message string `json:"message"`
	SentAt  int64  `json:"sent_at"`
}
