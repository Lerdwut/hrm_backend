package port

import (
	"hr_management/internal/core/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
	GetUserByUsername(id string) (*domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUser(id string) error
	FindByEmail(email string) (*domain.User, error)
	FindByGoogleID(googleID string) (*domain.User, error)
}

type UserService interface {
	Register(user *domain.User) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByGoogleID(googleID string) (*domain.User, error)
	CreateOAuthUser(oauthUser *domain.OAuthUser) (*domain.User, error)
	UpdateOAuthUser(user *domain.User, oauthUser *domain.OAuthUser) (*domain.User, error)
}
