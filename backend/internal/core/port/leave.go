package port

import (
	"hr_management/internal/core/domain"
)

type LeaveRepository interface {
	Create(leave *domain.Leave) error
	GetAll() ([]domain.Leave, error)
	GetByID(id uint) (*domain.Leave, error)
	Update(id uint, status domain.LeaveStatus) error
}