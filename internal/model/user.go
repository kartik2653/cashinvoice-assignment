package model

import (
	custom_errors "cashinvoice-assignment/internal/errors"

	"gorm.io/gorm"
)

type Role string

const (
	RoleUser  string = "user"
	RoleAdmin string = "admin"
)

type User struct {
	BaseModel
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"size:20;default:user"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Role != RoleUser && u.Role != RoleAdmin {
		return custom_errors.ErrInvalidRoleValue
	}
	return nil
}
