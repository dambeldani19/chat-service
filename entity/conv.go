package entity

import "time"

type Conversation struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
}

type Message struct {
	ID             uint      `gorm:"primaryKey"`
	ConversationID uint      `gorm:"not null"`                          // ForeignKey to Conversations
	SenderID       uint      `gorm:"not null"`                          // ForeignKey to Users
	Sender         User      `gorm:"foreignKey:SenderID;references:ID"` // Establish relationship with User
	Content        string    `gorm:"type:text"`
	SentAt         time.Time `gorm:"autoCreateTime"`
}
