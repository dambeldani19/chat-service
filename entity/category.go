package entity

import "time"

type Category struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type ArtikelCategory struct {
	ArtikelID  int64 `json:"artikel_id"`
	CategoryID int64 `json:"category_id"`
}
