package service

import (
	"chat-service/dto"
	"chat-service/repository"
)

type ArtikelService interface {
	GetListArtikel(categoryID int64) ([]dto.ArtikelRes, error)
}

type artikelService struct {
	repo repository.ArtikelRepo
}

func NewArtikelService(r repository.ArtikelRepo) *artikelService {
	return &artikelService{
		repo: r,
	}
}

func (s *artikelService) GetListArtikel(categoryID int64) ([]dto.ArtikelRes, error) {

	// var resp []dto.ArtikelRes
	list, err := s.repo.GetListArtikel(categoryID)
	if err != nil {
		return nil, err
	}

	// for _, a := range list {
	// 	resp = append(resp, dto.ArtikelRes{
	// 		ID:        a.ID,
	// 		CreatorID: a.CreatorID,
	// 		Creator: dto.UserResponse{
	// 			ID:        a.Creator.ID,
	// 			UserName:  a.Creator.Username,
	// 			Email:     a.Creator.Email,
	// 			CreatedAt: a.Creator.CreatedAt,
	// 		},
	// 		Title:        a.Title,
	// 		Content:      a.Content,
	// 		CreatedAt:    a.CreatedAt,
	// 		TotalLike:    a.TotalLike,
	// 		TotalComment: a.TotalComment,
	// 	})
	// }

	return list, err
}
