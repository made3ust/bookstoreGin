package handlers

import (
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var authors = []models.Author{{ID: 1, Name: "J.R.R. Tolkien"}}
var nextAuthorID = 2

func GetAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, authors)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err == nil {
		author.ID = nextAuthorID
		nextAuthorID++
		authors = append(authors, author)
		c.JSON(http.StatusCreated, author)
	}
}
