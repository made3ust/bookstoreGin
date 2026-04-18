package database

import (
	"bookstore/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host := "localhost"
	user := "postgres"
	password := "postgres"
	dbname := "bookstore"
	port := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Postgres: %v", err))
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Author{},
		&models.Category{},
		&models.Book{},
		&models.FavoriteBook{},
	)

	if err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}

	DB = db
	fmt.Println("Successfully connected to PostgreSQL")
}
