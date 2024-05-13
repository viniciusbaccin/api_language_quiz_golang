package main

import (
	"api_golang_ia/controllers"
	"github.com/gin-gonic/gin"
)

func enableCORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // You should restrict this in production
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func main() {
	router := gin.Default()

	// Apply CORS middleware to all routes
	router.Use(enableCORS)

	homeController := controllers.HomeController{}
	router.GET("/", homeController.Index)

	wordsController := controllers.WordsController{}
	router.GET("/words", wordsController.Index)

	router.Run(":8082")
}
