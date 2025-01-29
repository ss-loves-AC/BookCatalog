package main

import (
	"book-catalog/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"}) // Only trust localhost

	// Define routes
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:id", handlers.GetBook)
	router.POST("/books", handlers.CreateBook)
	router.PUT("/books/:id", handlers.UpdateBook)
	router.DELETE("/books/:id", handlers.DeleteBook)

	// Run server
	router.Run(":8080")
}
