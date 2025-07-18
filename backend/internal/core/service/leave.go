package service

import (
	"hr_management/internal/core/domain"
	"hr_management/internal/core/port"
)

type LeaveService struct {
	repo port.LeaveRepository
}

func NewLeaveService(r port.LeaveRepository) *LeaveService {
	return &LeaveService{repo: r}
}

func (s *LeaveService) RequestLeave(l *domain.Leave) error {
	return s.repo.Create(l)
}

func (s *LeaveService) ApprovedLeave(id uint) error {
	return s.repo.Update(id, domain.Approved)
}

func (s *LeaveService) RejectedLeave(id uint) error {
	return s.repo.Update(id, domain.Rejected)
}

func (s *LeaveService) GetAllLeaves() ([]domain.Leave, error) {
	return s.repo.GetAll()
}
