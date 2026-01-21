package service

import (
	"errors"

	"cashinvoice-assignment/internal/model"
	"cashinvoice-assignment/internal/repository"
	"cashinvoice-assignment/internal/utils"

	custom_errors "cashinvoice-assignment/internal/errors"

	"gorm.io/gorm"
)

type AuthService interface {
	Register(name, email, password, role string) error
	Login(email, password string) (*model.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(name, email, password, role string) error {
	existingUser, err := s.userRepo.FindByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if existingUser != nil && err == nil {
		return custom_errors.ErrUserAlreadyExists
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Role:     role,
	}

	return s.userRepo.Create(user)
}

func (s *authService) Login(email, password string) (*model.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if !utils.CheckPassword(user.Password, password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
