package database

import (
	"github.com/extndr/todo-go/internal/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDB() (*gorm.DB, error) {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "todos.db"
	}

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate without timestamps
	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		return nil, err
	}

	log.Println("SQLite DB initialized with GORM")
	return db, nil
}
