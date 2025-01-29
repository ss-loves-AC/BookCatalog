package main

import (
	"BookCatalog/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:id", handlers.GetBook)
	router.POST("/books", handlers.CreateBook)
	router.PUT("/books/:id", handlers.UpdateBook)		
	router.DELETE("/books/:id", handlers.DeleteBook)
	
	router.Run(":8080")
}
