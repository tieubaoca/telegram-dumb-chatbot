package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // Update this import
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tieubaoca/telegram-dumb-chatbot/docs"
)

// @title Gin API with Swagger
// @version 1.0
// @description This is a sample server for a Gin-based API.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()

	// Define Swagger API documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
