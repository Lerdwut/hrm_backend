package port

import (
	"hr_management/internal/core/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetUserByID(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByGoogleID(googleID string) (*domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUser(id string) error
}

type UserService interface {
	Register(user *domain.User) (*domain.User, error)
	ListUsers([]*domain.User, error)
}
