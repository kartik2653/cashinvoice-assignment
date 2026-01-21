package repository

import (
	"cashinvoice-assignment/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(todo *model.Todo) error
	GetByID(id uuid.UUID) (*model.Todo, error)
	GetByUser(userID uint) ([]model.Todo, error)
	GetByUserPaginated(userID uint, page, limit int, status string) ([]model.Todo, int64, error)
	Update(todo *model.Todo) error
	Delete(todo *model.Todo) error
}

type todoRepo struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepo{db: db}
}

func (r *todoRepo) Create(todo *model.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepo) GetByID(id uuid.UUID) (*model.Todo, error) {
	var todo model.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *todoRepo) GetByUser(userID uint) ([]model.Todo, error) {
	var todos []model.Todo
	err := r.db.Where("user_id = ?", userID).Find(&todos).Error
	return todos, err
}

func (r *todoRepo) GetByUserPaginated(userID uint, page, limit int, status string) ([]model.Todo, int64, error) {
	var todos []model.Todo
	var total int64

	query := r.db
	if userID != 0 {
		query = query.Where("user_id = ?", userID)
	}

	// Apply status filter if provided
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Get total count
	if err := query.Model(&model.Todo{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&todos).Error
	return todos, total, err
}

func (r *todoRepo) Update(todo *model.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepo) Delete(todo *model.Todo) error {
	return r.db.Delete(todo).Error
}
