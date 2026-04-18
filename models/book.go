package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title      string  `json:"title" binding:"required"`
	AuthorID   uint    `json:"author_id" binding:"required"`
	CategoryID uint    `json:"category_id" binding:"required"`
	Price      float64 `json:"price" binding:"required,gt=0"`
}
