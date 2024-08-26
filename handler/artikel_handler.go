package handler

import (
	"chat-service/dto"
	"chat-service/helper"
	"chat-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type artikelHandler struct {
	service service.ArtikelService
}

func NewArtikelHandler(s service.ArtikelService) artikelHandler {
	return artikelHandler{
		service: s,
	}
}

func (h *artikelHandler) GetListArtikel(c *gin.Context) {
	var catint int64
	categoryIDStr := c.Query("category_id")

	if categoryIDStr != "" {
		id, _ := strconv.Atoi(categoryIDStr)
		catint = int64(id)
	}

	list, err := h.service.GetListArtikel(catint)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	res := helper.Response(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Get list artikel success",
		Data:       list,
	})

	c.JSON(http.StatusOK, res)
}
