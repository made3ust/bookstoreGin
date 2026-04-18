package handlers

import (
	"bookstore/database"
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	database.DB.Find(&categories)
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var cat models.Category
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	database.DB.Create(&cat)
	c.JSON(http.StatusCreated, cat)
}
