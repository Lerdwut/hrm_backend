package port

import (
	"hr_management/internal/core/domain"
)

// internal/core/port/oauth.go
type GoogleService interface {
	GetAuthURL(state string) string
	HandleCallback(code, state string) (*domain.User, error)
	ValidateState(state string) error
	GenerateState() string
}
