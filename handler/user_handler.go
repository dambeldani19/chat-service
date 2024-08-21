package handler

import (
	"chat-service/dto"
	"chat-service/errorhandler"
	"chat-service/helper"
	"chat-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// type UserHandler interface{}

type userHandler struct {
	service service.AuthService
}

func NewUserHandler(s service.AuthService) *userHandler {
	return &userHandler{
		service: s,
	}
}

func (h *userHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}
	user, err := h.service.Register(&register)
	if err != nil {
		errorhandler.HandlerError(c, &errorhandler.InternalServerError{Message: err.Error()})
		return
	}

	res := helper.Response(dto.ResponseParam{
		StatusCode: http.StatusCreated,
		Message:    "Register success",
		Data: dto.UserResponse{
			ID:        user.ID,
			UserName:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
	})

	c.JSON(http.StatusCreated, res)

}

func (h *userHandler) Login(c *gin.Context) {
	var login dto.LoginRequest

	if err := c.ShouldBindJSON(&login); err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.service.Login(&login)
	if err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	res := helper.Response(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success login",
		Data:       result,
	})

	c.JSON((http.StatusOK), res)
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	data, err := h.service.GetDetail(intID)
	if err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	res := helper.Response(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       data,
	})

	c.JSON((http.StatusOK), res)
}
