package port

import "hr_management/internal/core/domain"

type AuthService interface {
	Login(username, password string) (*domain.User, error)
}
