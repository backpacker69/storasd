package models

type Message struct {
	ID        string `json:"id"`
	SenderID  string `json:"senderId"`
	ReceiverID string `json:"receiverId"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}