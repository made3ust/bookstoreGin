package models

type Book struct {
	ID         int     `json:"id"`
	Title      string  `json:"title" binding:"required"`
	AuthorID   int     `json:"author_id" binding:"required"`
	CategoryID int     `json:"category_id" binding:"required"`
	Price      float64 `json:"price" binding:"required,gt=0"`
}
