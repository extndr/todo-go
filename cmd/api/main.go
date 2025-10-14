package main

import (
	"github.com/extndr/todo-go/internal/database"
	"github.com/extndr/todo-go/internal/handlers"
	"github.com/extndr/todo-go/internal/repository"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewSQLiteDB()
	if err != nil {
		log.Fatal(err)
	}

	todoRepo := repository.NewTodoRepository(db)
	todoHandler := handlers.NewTodoHandler(todoRepo)

	r := gin.Default()

	todos := r.Group("/todos")
	{
		todos.GET("/", todoHandler.GetTodos)
		todos.GET("/:id", todoHandler.GetTodo)
		todos.POST("/", todoHandler.CreateTodo)
		todos.PUT("/:id", todoHandler.UpdateTodo)
		todos.DELETE("/:id", todoHandler.DeleteTodo)
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
