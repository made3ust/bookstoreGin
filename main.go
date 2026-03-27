package main

import (
	"bookstore/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	bookRoutes := r.Group("/books")
	{
		bookRoutes.GET("", handlers.GetBooks)
		bookRoutes.POST("", handlers.CreateBook)
		bookRoutes.GET("/:id", handlers.GetBook)
		bookRoutes.PUT("/:id", handlers.UpdateBook)
		bookRoutes.DELETE("/:id", handlers.DeleteBook)
	}

	r.GET("/authors", handlers.GetAuthors)
	r.POST("/authors", handlers.CreateAuthor)

	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.CreateCategory)

	r.Run(":8080")
}
