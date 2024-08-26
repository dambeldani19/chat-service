package dto

import "time"

type ArtikelRes struct {
	ID           int64              `json:"id"`
	CreatorID    int64              `json:"creator_id"`
	Creator      UserResponse       `json:"creator"`
	Categories   []CategoryResponse `json:"categories"`
	Title        string             `json:"title"`
	Content      string             `json:"content"`
	TotalLike    int                `json:"total_like"`
	TotalComment int                `json:"total_comment"`
	CreatedAt    time.Time          `json:"created_at"`
}

type CategoryResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
