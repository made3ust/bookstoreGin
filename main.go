package main

import (
	"bookstore/database"
	"bookstore/handlers"
	"bookstore/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		bookRoutes := auth.Group("/books")
		{
			bookRoutes.GET("/favorites", handlers.GetFavorites)

			bookRoutes.GET("", handlers.GetBooks)
			bookRoutes.POST("", handlers.CreateBook)
			bookRoutes.GET("/:id", handlers.GetBook)
			bookRoutes.PUT("/:id", handlers.UpdateBook)
			bookRoutes.DELETE("/:id", handlers.DeleteBook)

			bookRoutes.PUT("/:id/favorites", handlers.AddFavorite)
			bookRoutes.DELETE("/:id/favorites", handlers.RemoveFavorite)
		}

		auth.GET("/authors", handlers.GetAuthors)
		auth.POST("/authors", handlers.CreateAuthor)
		auth.GET("/categories", handlers.GetCategories)
		auth.POST("/categories", handlers.CreateCategory)
	}

	r.Run(":8070")
}
