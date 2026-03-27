package handlers

import (
	"bookstore/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{
	{ID: 1, Title: "The Hobbit", AuthorID: 1, CategoryID: 1, Price: 15.99},
}
var nextBookID = 2

func GetBooks(c *gin.Context) {
	categoryIDStr := c.Query("category_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	filteredBooks := books

	if categoryIDStr != "" {
		catID, _ := strconv.Atoi(categoryIDStr)
		var temp []models.Book
		for _, b := range books {
			if b.CategoryID == catID {
				temp = append(temp, b)
			}
		}
		filteredBooks = temp
	}

	start := (page - 1) * limit
	end := start + limit
	if start > len(filteredBooks) {
		filteredBooks = []models.Book{}
	} else {
		if end > len(filteredBooks) {
			end = len(filteredBooks)
		}
		filteredBooks = filteredBooks[start:end]
	}

	c.JSON(http.StatusOK, filteredBooks)
}

func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, item := range books {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func CreateBook(c *gin.Context) {
	var newBook models.Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if newBook.Title == "" || newBook.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required and Price > 0"})
		return
	}

	newBook.ID = nextBookID
	nextBookID++
	books = append(books, newBook)

	c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedBook models.Book

	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for i, item := range books {
		if item.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
