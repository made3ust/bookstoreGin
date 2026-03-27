package handlers

import (
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var categories = []models.Category{
	{ID: 1, Name: "Fantasy"},
}
var nextCategoryID = 2

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var cat models.Category

	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if cat.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category name is required"})
		return
	}

	cat.ID = nextCategoryID
	nextCategoryID++
	categories = append(categories, cat)

	c.JSON(http.StatusCreated, cat)
}
