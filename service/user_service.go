package service

import (
	"chat-service/dto"
	"chat-service/entity"
	"chat-service/errorhandler"
	"chat-service/helper"
	"chat-service/repository"
	"time"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) (*entity.User, error)
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	GetDetail(id int) (*dto.UserResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) (*entity.User, error) {

	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return nil, &errorhandler.BadRequestError{Message: "email already registered"}
	}

	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	user := entity.User{
		Username:  req.UserName,
		Email:     req.Email,
		Password:  passwordHash,
		CreatedAt: time.Now(),
	}

	_, err = s.repository.Register(&user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	return &user, nil

}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data *dto.LoginResponse

	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "wrong email or password"}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.NotFoundError{Message: "wrong email or password"}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = &dto.LoginResponse{
		ID:       user.ID,
		UserName: user.Username,
		Email:    user.Email,
		Token:    token,
	}
	return data, nil
}

func (s *authService) GetDetail(id int) (*dto.UserResponse, error) {

	user, err := s.repository.GetDetailUser(id)
	if err != nil {
		return nil, &errorhandler.BadRequestError{Message: "data user tidak di temukan"}
	}

	res := dto.UserResponse{
		ID:        user.ID,
		UserName:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return &res, nil
}
