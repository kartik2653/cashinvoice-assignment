package service

import (
	custom_errors "cashinvoice-assignment/internal/errors"
	"cashinvoice-assignment/internal/model"
	"cashinvoice-assignment/internal/repository"
	"cashinvoice-assignment/internal/utils"

	"github.com/google/uuid"
)

type TodoService interface {
	CreateTodo(title string, description string, status model.TodoStatus, userID uint) (*model.Todo, error)
	GetTodos(userID uint, role string, page, limit int, status string) ([]model.Todo, int64, error)
	UpdateTodo(id uuid.UUID, title, description string, status model.TodoStatus, userID uint) (*model.Todo, error)
	DeleteTodo(id uuid.UUID, userID uint) error
}

type todoService struct {
	todoRepo repository.TodoRepository
	worker   *utils.AutoCompleteWorker
}

func NewTodoService(todoRepo repository.TodoRepository, worker *utils.AutoCompleteWorker) TodoService {
	return &todoService{
		todoRepo: todoRepo,
		worker:   worker,
	}
}

func (s *todoService) CreateTodo(title string, description string, status model.TodoStatus, userID uint) (*model.Todo, error) {
	todo := &model.Todo{
		Title:       title,
		Description: description,
		Status:      status,
		UserID:      userID,
	}
	s.worker.TaskChan <- todo.ID
	return todo, s.todoRepo.Create(todo)
}

func (s *todoService) GetTodos(userID uint, role string, page, limit int, status string) ([]model.Todo, int64, error) {
	if role == "admin" {
		// Admin sees all todos
		return s.todoRepo.GetByUserPaginated(0, page, limit, status) // 0 means fetch all
	}
	return s.todoRepo.GetByUserPaginated(userID, page, limit, status)
}

func (s *todoService) UpdateTodo(id uuid.UUID, title, description string, status model.TodoStatus, userID uint) (*model.Todo, error) {
	todo, err := s.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if todo.UserID != userID {
		return nil, custom_errors.ErrUnauthorized
	}
	todo.Title = title
	todo.Description = description
	todo.Status = status

	return todo, s.todoRepo.Update(todo)
}

func (s *todoService) DeleteTodo(id uuid.UUID, userID uint) error {
	todo, err := s.todoRepo.GetByID(id)
	if err != nil {
		return err
	}

	if todo.UserID != userID {
		return custom_errors.ErrUnauthorized
	}

	return s.todoRepo.Delete(todo)
}
