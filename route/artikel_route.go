package route

import (
	"chat-service/config"
	"chat-service/handler"
	"chat-service/repository"
	"chat-service/service"

	"github.com/gin-gonic/gin"
)

func ArtikelRouter(api *gin.RouterGroup) {
	repo := repository.NewArtikelRepo(config.DB)
	srv := service.NewArtikelService(repo)
	h := handler.NewArtikelHandler(srv)

	r := api.Group("/artikel")

	r.GET("/", h.GetListArtikel)
}
