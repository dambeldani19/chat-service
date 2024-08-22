package test

import (
	"chat-service/entity"
	"chat-service/middleware"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"chat-service/handler"
	"chat-service/repository"
	"chat-service/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// AuthMiddleware extracts the userID from the Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		userID := parts[1]
		c.Set("userID", userID)
		c.Next()
	}
}

// SetupRouter for testing
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Setup in-memory database
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&entity.User{}, &entity.Message{}, &entity.Conversation{})

	// Initialize repositories
	userRepo := repository.NewAuthRepository(db)
	conversationRepo := repository.NewconvRepository(db)

	// Initialize services
	userService := service.NewAuthService(userRepo)
	messageService := service.NewConversationService(conversationRepo, userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	messageHandler := handler.NewConvHandler(messageService)

	// Define routes without middleware
	r.POST("/users", userHandler.Register)
	r.POST("/login", userHandler.Login) // No AuthMiddleware here

	// Define routes with middleware
	authRoutes := r.Group("/")
	authRoutes.Use(middleware.JWTMiddleware(db))
	// authRoutes.Use(AuthMiddleware())
	{
		authRoutes.POST("/conversations", messageHandler.CreateConversation)
		authRoutes.POST("/conversations/:conversationId/messages", messageHandler.SendMessage)
		authRoutes.GET("/conversations/:conversationId/messages", messageHandler.GetMessagesByConversationID)
	}

	return r
}

func TestCreateUser(t *testing.T) {
	r := SetupRouter()

	// Sample user payload
	payload := `{"username":"johndoe","email":"johndoe@example.com","password":"password123"}`

	req, _ := http.NewRequest("POST", "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Validate response
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"username":"johndoe"`)
	assert.Contains(t, w.Body.String(), `"email":"johndoe@example.com"`)
}

func TestLogin(t *testing.T) {
	r := SetupRouter()

	// Create a user with ID 1
	userPayload := `{"username":"johndoe","email":"johndoe@example.com","password":"password123"}`
	userReq, _ := http.NewRequest("POST", "/users", strings.NewReader(userPayload))
	userReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, userReq)

	// Test login
	loginPayload := `{"email":"johndoe@example.com","password":"password123"}`
	loginReq, _ := http.NewRequest("POST", "/login", strings.NewReader(loginPayload))
	loginReq.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, loginReq)

	// Validate response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"token"`)
}

func TestCreateConversation(t *testing.T) {
	r := SetupRouter()

	// Create a user with ID 1
	userPayload := `{"username":"johndoe","email":"johndoe@example.com","password":"password123"}`
	userReq, _ := http.NewRequest("POST", "/users", strings.NewReader(userPayload))
	userReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, userReq)

	// Login and get token
	loginPayload := `{"email":"johndoe@example.com","password":"password123"}`
	loginReq, _ := http.NewRequest("POST", "/login", strings.NewReader(loginPayload))
	loginReq.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, loginReq)

	// Print the login response for debugging
	t.Log("Login response:", w.Body.String())

	// Extract token from response
	var loginResponse struct {
		Code    int    `json:"code"`
		Status  string `json:"status"`
		Message string `json:"message"`
		Data    struct {
			Token string `json:"token"`
		} `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &loginResponse); err != nil {
		t.Fatalf("could not unmarshal login response: %v", err)
	}
	token := loginResponse.Data.Token
	if token == "" {
		t.Fatal("token not found in login response")
	}

	// Create a conversation (auth needed here)
	convoPayload := `{"title":"Test Conversation"}`
	convoReq, _ := http.NewRequest("POST", "/conversations", strings.NewReader(convoPayload))
	convoReq.Header.Set("Content-Type", "application/json")
	convoReq.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, convoReq)

	// Sample message payload
	payload := `{"sender_id":1,"content":"Hello, how are you?"}`

	req, _ := http.NewRequest("POST", "/conversations/1/messages", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Validate response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"content":"Hello, how are you?"`)
}
