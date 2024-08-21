package route

import (
	"chat-service/config"
	"chat-service/handler"
	"chat-service/middleware"
	"chat-service/repository"
	"chat-service/service"

	"github.com/gin-gonic/gin"
)

func ConvRouter(api *gin.RouterGroup) {
	convRepo := repository.NewconvRepository(config.DB)
	userRepo := repository.NewAuthRepository(config.DB)
	convService := service.NewConversationService(convRepo, userRepo)
	h := handler.NewConvHandler(convService)

	r := api.Group("/conversations")
	r.Use(middleware.JWTMiddleware(config.DB))

	r.POST("/", h.CreateConversation)
	r.POST("/:conversationId/messages", h.SendMessage)
	r.GET("/:conversationId/messages", h.GetMessagesByConversationID)
}
