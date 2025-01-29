package handlers

import (
    "book-catalog/models"
    "book-catalog/storage"
    "net/http"
    "strconv"
    "log"

    "github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
    log.Println("Fetching all books")
    books := make([]models.Book, 0, len(storage.Books))
    for _, book := range storage.Books {
        books = append(books, book)
    }
    log.Printf("Total books retrieved: %d\n", len(books))
    c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
    id := c.Param("id")
    log.Printf("Fetching book with ID: %s\n", id)
    if book, exists := storage.Books[id]; exists {
        log.Printf("Book found: %+v\n", book)
        c.JSON(http.StatusOK, book)
        return
    }
    log.Println("Book not found")
    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func CreateBook(c *gin.Context) {
    log.Println("Creating a new book")
    var newBook models.Book
    if err := c.ShouldBindJSON(&newBook); err != nil {
        log.Printf("Error parsing request body: %v\n", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    storage.LastID++
    newBook.ID = strconv.Itoa(storage.LastID)
    storage.Books[newBook.ID] = newBook
    log.Printf("Book created successfully: %+v\n", newBook)
    c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    log.Printf("Updating book with ID: %s\n", id)

    var updatedBook models.Book
    if err := c.ShouldBindJSON(&updatedBook); err != nil {
        log.Printf("Error parsing request body: %v\n", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if _, exists := storage.Books[id]; exists {
        updatedBook.ID = id
        storage.Books[id] = updatedBook
        log.Printf("Book updated successfully: %+v\n", updatedBook)
        c.JSON(http.StatusOK, updatedBook)
        return
    }

    log.Println("Book not found for update")
    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    log.Printf("Deleting book with ID: %s\n", id)
    if _, exists := storage.Books[id]; exists {
        delete(storage.Books, id)
        log.Println("Book deleted successfully")
        c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
        return
    }
    log.Println("Book not found for deletion")
    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
