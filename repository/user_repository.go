package repository

import (
	"chat-service/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	EmailExist(email string) bool
	Register(user *entity.User) (int, error)
	GetUserByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	GetDetailUser(id int) (*entity.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(user *entity.User) (int, error) {

	err := r.db.Create(&user).Error

	return user.ID, err
}

func (r *authRepository) EmailExist(email string) bool {
	var user entity.User

	err := r.db.First(&user, "email = ?", email).Error

	return err == nil
}

func (r *authRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.db.First(&user, "email = ?", email).Error

	return &user, err
}

func (r *authRepository) Update(user *entity.User) error {
	err := r.db.Model(&user).Updates(&user).Error

	return err
}

func (r *authRepository) GetDetailUser(id int) (*entity.User, error) {
	var user *entity.User

	err := r.db.First(&user, "id = ?", id).Error

	return user, err
}
