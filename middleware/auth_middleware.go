package middleware

import (
	"chat-service/errorhandler"
	"chat-service/helper"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JWTMiddleware(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		fmt.Println(tokenString)

		if tokenString == "" {
			errorhandler.HandlerError(c, &errorhandler.UnathorizedError{Message: "Unauthorize"})
			c.Abort()
			return
		}

		// Memeriksa apakah header memiliki skema Bearer
		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		// Mengambil token dari header dengan menghilangkan "Bearer "
		token := strings.TrimPrefix(tokenString, "Bearer ")

		userId, err := helper.ValidateToken(token)
		if err != nil {
			errorhandler.HandlerError(c, &errorhandler.UnathorizedError{Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", *userId)
		c.Next()

	}
}
