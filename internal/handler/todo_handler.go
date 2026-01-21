package handler

import (
	"cashinvoice-assignment/internal/model"
	"cashinvoice-assignment/internal/service"
	"errors"

	custom_errors "cashinvoice-assignment/internal/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TodoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

type CreateTodoRequest struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Status      model.TodoStatus `json:"status"`
}

type UpdateTodoRequest struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Status      model.TodoStatus `json:"status"`
}

func (h *TodoHandler) Create(c *fiber.Ctx) error {
	var req CreateTodoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	userID := c.Locals("user_id").(uint)
	todo, err := h.todoService.CreateTodo(req.Title, req.Description, req.Status, userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(todo)
}

func (h *TodoHandler) List(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}

	role, ok := c.Locals("role").(string)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Validate status filter
	status := c.Query("status", "")
	switch status {
	case "", "pending", "in_progress", "completed":
	default:
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid status. must be one of: pending, in_progress, completed",
		})
	}

	todos, total, err := h.todoService.GetTodos(userID, role, page, limit, status)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	return c.JSON(fiber.Map{
		"data": todos,
		"meta": fiber.Map{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}

func (h *TodoHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	var req UpdateTodoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	userID := c.Locals("user_id").(uint)

	todo, err := h.todoService.UpdateTodo(id, req.Title, req.Description, req.Status, userID)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUnauthorized) {
			return c.Status(403).JSON(fiber.Map{"error": "forbidden"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(todo)
}

func (h *TodoHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	userID := c.Locals("user_id").(uint)

	err = h.todoService.DeleteTodo(id, userID)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUnauthorized) {
			return c.Status(403).JSON(fiber.Map{"error": "forbidden"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(204)
}
