package route

import (
	"chat-service/config"
	"chat-service/handler"
	"chat-service/repository"
	"chat-service/service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewUserHandler(authService)

	api.POST("/users", authHandler.Register)
	api.GET("/users/:id", authHandler.GetUserByID)
	api.POST("/login", authHandler.Login)
}
