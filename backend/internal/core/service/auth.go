package service

import (
	"hr_management/internal/core/domain"
	"hr_management/internal/core/port"
)

type AuthService struct {
	repo port.UserRepository
}

func NewAuthService(userRepository port.UserRepository) *AuthService {
	return &AuthService{repo: userRepository}
}

func (as *AuthService) Login(username, password string) (*domain.User, error) {
	repoUser, err := as.repo.GetUserByUsername(username)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, domain.ErrDataNotFound
		}
		return nil, domain.ErrInternal
	}

	if repoUser == nil {
		return nil, domain.ErrDataNotFound
	}

	if repoUser.Password != password {
		return nil, domain.ErrInvalidCredentials
	}

	return repoUser, nil
}
