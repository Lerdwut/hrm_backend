package port

import (
	"hr_management/internal/core/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
	FindByUsername(id string) (*domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUser(id string) error
	FindByEmail(email string) (*domain.User, error)
	FindByGoogleID(googleID string) (*domain.User, error)
}

type UserService interface {
	Register(user *domain.User) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByGoogleID(googleID string) (*domain.User, error)
	CreateGoogleUser(oauthUser *domain.GoogleUser) (*domain.User, error)
	UpdateGoogleUser(user *domain.User, oauthUser *domain.GoogleUser) (*domain.User, error)
}
