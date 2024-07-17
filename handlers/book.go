package handlers

import (
	"go-gin-crud/database"
	"go-gin-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	c.IndentedJSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	var book models.Book
	result := database.DB.First(&book, bookID)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	database.DB.Create(&newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	var updatedBook models.Book
	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}

	var book models.Book
	result := database.DB.First(&book, bookID)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	database.DB.Save(&book)

	c.IndentedJSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	var book models.Book
	result := database.DB.First(&book, bookID)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	database.DB.Delete(&book)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
