package handlers

import (
    "book-catalog/models"
    "book-catalog/storage"
    "net/http"
	"strconv"
    "github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
    books := make([]models.Book, 0, len(storage.Books))
    for _, book := range storage.Books {
        books = append(books, book)
    }
    c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
    id := c.Param("id")
    if book, exists := storage.Books[id]; exists {
        c.JSON(http.StatusOK, book)
        return
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func CreateBook(c *gin.Context) {
    var newBook models.Book
    if err := c.ShouldBindJSON(&newBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    storage.LastID++
    newBook.ID = strconv.Itoa(storage.LastID)
    storage.Books[newBook.ID] = newBook
    c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var updatedBook models.Book
    if err := c.ShouldBindJSON(&updatedBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if _, exists := storage.Books[id]; exists {
        updatedBook.ID = id
        storage.Books[id] = updatedBook
        c.JSON(http.StatusOK, updatedBook)
        return
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    if _, exists := storage.Books[id]; exists {
        delete(storage.Books, id)
        c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
        return
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
