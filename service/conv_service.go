package service

import (
	"chat-service/entity"
	"chat-service/errorhandler"
	"chat-service/repository"
	"fmt"
	"time"
)

type ConversationService interface {
	CreateConversation() (*entity.Conversation, error)
	SendMessage(conversationID, senderID uint, content string) (*entity.Message, error)
	GetMessagesByConversationID(conversationID, userLogin uint) ([]entity.Message, error)
}

type conversationService struct {
	repository repository.ConvRepository
	userRepo   repository.AuthRepository
}

func NewConversationService(r repository.ConvRepository, us repository.AuthRepository) *conversationService {
	return &conversationService{
		repository: r,
		userRepo:   us,
	}
}

func (s *conversationService) CreateConversation() (*entity.Conversation, error) {
	conversation := &entity.Conversation{}
	err := s.repository.CreateConversation(conversation)
	if err != nil {
		return nil, err
	}
	return conversation, nil
}

func (s *conversationService) SendMessage(conversationID, senderID uint, content string) (*entity.Message, error) {

	// check sender id exist
	_, err := s.userRepo.GetDetailUser(int(senderID))
	if err != nil {
		return nil, &errorhandler.BadRequestError{Message: fmt.Sprintf("sender id %d tidak di temukan", senderID)}
	}

	// check conversation id exist
	_, err = s.repository.GetConversationByID(conversationID)
	if err != nil {
		return nil, &errorhandler.BadRequestError{Message: fmt.Sprintf("data conversation id %d tidak di temukan", conversationID)}
	}

	message := &entity.Message{
		ConversationID: conversationID,
		SenderID:       senderID,
		Content:        content,
		SentAt:         time.Now(),
	}
	err = s.repository.CreateMessage(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (s *conversationService) GetMessagesByConversationID(conversationID, userLogin uint) ([]entity.Message, error) {

	// check conversation exists
	_, err := s.repository.GetConversationByID(uint(conversationID))
	if err != nil {
		return nil, &errorhandler.BadRequestError{Message: fmt.Sprintf("data conversation id %d tidak di temukan", conversationID)}
	}

	return s.repository.GetMessagesByConversationID(conversationID, userLogin)
}
