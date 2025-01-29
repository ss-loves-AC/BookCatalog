package handlers


import (
	"github.com/gin-gonic/gin"
	"BookCatalog/models"
	"BookCatalog/storage"
)

func GetBooks(c *gin.Context) {
	for _, book := range storage.Books {
		c.JSON(200, book)
	}
}

func GetBook(c *gin.Context) {
	// Get book by ID
	bookID := c.Param("id")
	if storage.Books[bookID] == nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}	
	c.JSON(200, storage.Books[bookID])
}

func CreateBook(c *gin.Context) {
	// Create new book
}

func UpdateBook(c *gin.Context) {
	// Update existing book
}

func DeleteBook(c *gin.Context) {
	// Delete book
}
