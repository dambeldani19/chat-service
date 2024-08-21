package main

import (
	"chat-service/config"
	"chat-service/route"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()
	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	route.AuthRouter(api)
	route.ConvRouter(api)

	r.Run(fmt.Sprintf(":%v", config.ENV.PORT)) // listen and serve on 0.0.0.0:8080
}
