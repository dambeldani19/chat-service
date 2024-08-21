package handler

import (
	"chat-service/dto"
	"chat-service/errorhandler"
	"chat-service/helper"
	"chat-service/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type convHandler struct {
	service service.ConversationService
}

func NewConvHandler(s service.ConversationService) *convHandler {
	return &convHandler{
		service: s,
	}
}

func (h *convHandler) CreateConversation(c *gin.Context) {
	conversation, err := h.service.CreateConversation()
	if err != nil {
		errorhandler.HandlerError(c, &errorhandler.InternalServerError{Message: err.Error()})
		return
	}

	res := helper.Response(dto.ResponseParam{
		StatusCode: http.StatusCreated,
		Message:    "Create conversation success",
		Data: dto.ConversationRes{
			ID:        int(conversation.ID),
			CreatedAt: conversation.CreatedAt,
		},
	})

	c.JSON(http.StatusOK, res)
}

func (h *convHandler) SendMessage(c *gin.Context) {
	user, _ := c.Get("userID")
	userInt := user.(int)

	var req *dto.SendMessageReq

	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	//validate request sender_id with login userID
	if userInt != int(req.SenderID) {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{Message: fmt.Sprintf("sender id %d tidak sesuai dengan user login", req.SenderID)})
		return
	}

	conversationID, _ := strconv.ParseUint(c.Param("conversationId"), 10, 32)
	conv, err := h.service.SendMessage(uint(conversationID), req.SenderID, req.Content)
	if err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	res := helper.Response(dto.ResponseParam{
		StatusCode: http.StatusCreated,
		Message:    "Create conversation success",
		Data: dto.ConversationSendMessageRes{
			ID:             int(conv.ID),
			ConversationID: int64(conv.ConversationID),
			SenderID:       int64(conv.SenderID),
			Content:        conv.Content,
			SendAt:         conv.SentAt,
		},
	})

	c.JSON(http.StatusOK, res)
}

func (h *convHandler) GetMessagesByConversationID(c *gin.Context) {

	user, _ := c.Get("userID")
	userInt := user.(int)

	var datasConv []dto.ConversationSendMessageRes

	// Parse conversationId from the URL parameter
	conversationID, err := strconv.ParseUint(c.Param("conversationId"), 10, 32)
	if err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	// Get all messages for the given conversation ID
	messages, err := h.service.GetMessagesByConversationID(uint(conversationID), uint(userInt))
	if err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	for _, m := range messages {
		datasConv = append(datasConv, dto.ConversationSendMessageRes{
			ID:             int(m.ID),
			ConversationID: int64(m.ConversationID),
			SenderID:       int64(m.SenderID),
			Content:        m.Content,
			SendAt:         m.SentAt,
		})
	}

	res := helper.Response(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Get list conversation success",
		Data:       datasConv,
	})

	c.JSON(http.StatusOK, res)
}
