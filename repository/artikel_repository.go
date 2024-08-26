package repository

import (
	"chat-service/dto"
	"chat-service/entity"

	"gorm.io/gorm"
)

type ArtikelRepo interface {
	GetListArtikel(categoryID int64) ([]dto.ArtikelRes, error)
}

type artikelRepo struct {
	db *gorm.DB
}

func NewArtikelRepo(db *gorm.DB) *artikelRepo {
	return &artikelRepo{
		db: db,
	}
}
func (r *artikelRepo) GetListArtikel(categoryID int64) ([]dto.ArtikelRes, error) {
	var arrArtikel []entity.Artikel
	var resp []dto.ArtikelRes

	query := r.db.
		Preload("Creator").
		Preload("Categories"). // Preload Categories
		Joins("LEFT JOIN artikel_categoris ac ON artikels.id = ac.artikel_id")

	if categoryID != 0 {
		query = query.Where("ac.category_id = ?", categoryID)
	}

	err := query.Find(&arrArtikel).Error
	if err != nil {
		return nil, err
	}

	for i, a := range arrArtikel {

		var commentCount int64
		var emojiCount int64

		// Hitung total komentar
		err = r.db.Model(&entity.ArtikelComment{}).
			Where("artikel_id = ?", arrArtikel[i].ID).
			Count(&commentCount).Error
		if err != nil {
			return nil, err
		}

		// Hitung total emoji
		err = r.db.Model(&entity.ArtikelEmoji{}).
			Where("artikel_id = ?", arrArtikel[i].ID).
			Count(&emojiCount).Error
		if err != nil {
			return nil, err
		}

		// Map categories
		var categories []dto.CategoryResponse
		for _, cat := range a.Categories {
			categories = append(categories, dto.CategoryResponse{
				ID:   cat.ID,
				Name: cat.Name,
			})
		}

		resp = append(resp, dto.ArtikelRes{
			ID:        a.ID,
			CreatorID: a.CreatorID,
			Creator: dto.UserResponse{
				ID:        a.Creator.ID,
				UserName:  a.Creator.Username,
				Email:     a.Creator.Email,
				CreatedAt: a.Creator.CreatedAt,
			},
			Title:        a.Title,
			Content:      a.Content,
			CreatedAt:    a.CreatedAt,
			TotalLike:    int(emojiCount),
			TotalComment: int(commentCount),
			Categories:   categories, // Tambahkan data kategori di sini
		})
	}

	return resp, nil
}
