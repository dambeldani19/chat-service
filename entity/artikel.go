package entity

import "time"

type Artikel struct {
	ID         int64      `json:"id"`
	CreatorID  int64      `json:"creator_id"`
	Creator    User       `gorm:"foreignKey:CreatorID;references:ID"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	Categories []Category `gorm:"many2many:artikel_categoris;"`
	CreatedAt  time.Time  `json:"created_at"`
}

type ArtikelComment struct {
	ID        int64     `json:"id"`
	ArtikelID int64     `json:"artikel_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

type ArtikelEmoji struct {
	ID        int64     `json:"id"`
	ArtikelID int64     `json:"artikel_id"`
	EmojiType string    `json:"emoji_type"`
	CreatedAt time.Time `json:"created_at"`
}
