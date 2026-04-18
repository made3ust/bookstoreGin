package handlers

import (
	"bookstore/database"
	"bookstore/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFavorites(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var favorites []models.FavoriteBook
	database.DB.Preload("Book").
		Where("user_id = ?", userID).
		Limit(limit).Offset(offset).
		Find(&favorites)

	var books []models.Book
	for _, f := range favorites {
		books = append(books, f.Book)
	}

	c.JSON(http.StatusOK, books)
}

func AddFavorite(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	bookID, _ := strconv.Atoi(c.Param("id"))

	favorite := models.FavoriteBook{
		UserID: userID,
		BookID: uint(bookID),
	}

	if err := database.DB.Create(&favorite).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not add to favorites"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Added to favorites"})
}

func RemoveFavorite(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	bookID, _ := strconv.Atoi(c.Param("id"))

	database.DB.Where("user_id = ? AND book_id = ?", userID, bookID).Delete(&models.FavoriteBook{})
	c.Status(http.StatusNoContent)
}
