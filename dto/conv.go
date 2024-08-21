package dto

import "time"

type ConversationRes struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type SendMessageReq struct {
	SenderID uint   `json:"sender_id"`
	Content  string `json:"content"`
}

type ConversationSendMessageRes struct {
	ID             int       `json:"id"`
	ConversationID int64     `json:"conversation_id"`
	SenderID       int64     `json:"sender_id"`
	Content        string    `json:"content"`
	SendAt         time.Time `json:"sent_at"`
}
