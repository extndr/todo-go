package repository

import (
	"github.com/extndr/todo-go/internal/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	if err := r.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) GetByID(id uint) (*models.Todo, error) {
	var t models.Todo
	if err := r.DB.First(&t, id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TodoRepository) Create(todo *models.Todo) error {
	return r.DB.Create(todo).Error
}

func (r *TodoRepository) Update(todo *models.Todo) error {
	return r.DB.Save(todo).Error
}

func (r *TodoRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Todo{}, id).Error
}
