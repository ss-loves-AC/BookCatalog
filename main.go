package main

import (
	"book-catalog/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"}) // Only trust localhost

	// Serve static files (HTML, CSS, JS) from the "public" directory
    router.Static("/assets", "./public")

    // Serve the index.html file when accessing root path "/"
    router.GET("/", func(c *gin.Context) {
        c.File("./public/index.html")
    })

	// Define routes
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:id", handlers.GetBook)
	router.POST("/books", handlers.CreateBook)
	router.PUT("/books/:id", handlers.UpdateBook)
	router.DELETE("/books/:id", handlers.DeleteBook)

	// Run server
	router.Run(":8080")
}
