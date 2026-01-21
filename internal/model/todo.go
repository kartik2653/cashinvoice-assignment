package model

import (
	custom_errors "cashinvoice-assignment/internal/errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoStatus string

const (
	Pending    TodoStatus = "pending"
	InProgress TodoStatus = "in_progress"
	Completed  TodoStatus = "completed"
)

type Todo struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Title       string     `gorm:"size:255;not null"`
	Description string     `gorm:"type:text"`
	Status      TodoStatus `gorm:"type:varchar(20);default:'pending'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      uint `gorm:"not null"`
}

func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()

	if err := t.validateStatus(); err != nil {
		return err
	}
	return nil
}

func (t *Todo) BeforeUpdate(tx *gorm.DB) (err error) {
	return t.validateStatus()
}

func (t *Todo) validateStatus() error {
	switch t.Status {
	case Pending, InProgress, Completed:
		return nil
	default:
		return custom_errors.ErrInvalidTaskStatus
	}
}
