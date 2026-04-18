package handlers

import (
	"bookstore/database"
	"bookstore/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	categoryIDStr := c.Query("category_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	query := database.DB

	if categoryIDStr != "" {
		catID, _ := strconv.Atoi(categoryIDStr)
		query = query.Where("category_id = ?", catID)
	}

	if err := query.Limit(limit).Offset(offset).Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook models.Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := database.DB.Create(&newBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create book"})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	book.ID = uint(id)
	database.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	result := database.DB.Delete(&models.Book{}, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
