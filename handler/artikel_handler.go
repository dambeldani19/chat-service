package handler

import (
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

	c.JSON(http.StatusOK, list)
}
