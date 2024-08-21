package repository

import (
	"chat-service/entity"
	"errors"

	"gorm.io/gorm"
)

type ConvRepository interface {
	CreateConversation(conversation *entity.Conversation) error
	GetConversationByID(id uint) (*entity.Conversation, error)
	CreateMessage(message *entity.Message) error
	GetMessagesByConversationID(conversationID, userLogin uint) ([]entity.Message, error)
}

type convRepository struct {
	db *gorm.DB
}

func NewconvRepository(db *gorm.DB) *convRepository {
	return &convRepository{
		db: db,
	}
}

func (r *convRepository) CreateConversation(conversation *entity.Conversation) error {
	return r.db.Create(conversation).Error
}

func (r *convRepository) GetConversationByID(id uint) (*entity.Conversation, error) {
	var conversation entity.Conversation
	err := r.db.First(&conversation, id).Error
	return &conversation, err
}

func (r *convRepository) CreateMessage(message *entity.Message) error {
	return r.db.Create(message).Error
}

func (r *convRepository) GetMessagesByConversationID(conversationID, userLogin uint) ([]entity.Message, error) {
	var messages []entity.Message

	var count int64
	err := r.db.Model(&entity.Message{}).
		Where("conversation_id = ? AND sender_id = ?", conversationID, userLogin).
		Count(&count).Error
	if err != nil {
		return messages, err
	}

	if count == 0 {
		return messages, errors.New("data conversation tidak ada")
	}

	err = r.db.Where("conversation_id = ?", conversationID).Find(&messages).Error
	return messages, err
}
