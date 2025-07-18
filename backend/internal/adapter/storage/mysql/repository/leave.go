package repository

import (
	"gorm.io/gorm"
	"hr_management/internal/core/domain"
)

type LeaveRepo struct {
	db *gorm.DB
}

func NewGormLeaveRepo(db *gorm.DB) *LeaveRepo {
	return &LeaveRepo{db: db}
}
func (r *LeaveRepo) Create(l *domain.Leave) error {
	return r.db.Create(l).Error
}

func (r *LeaveRepo) GetAll() ([]domain.Leave, error) {
	var leaves []domain.Leave
	err := r.db.Find(&leaves).Error
	return leaves, err
}

func (r *LeaveRepo) GetByID(id uint) (*domain.Leave, error) {
	var leave domain.Leave
	err := r.db.First(&leave, id).Error
	if err != nil {
		return nil, err
	}
	return &leave, nil
}

func (r *LeaveRepo) Update(id uint, status domain.LeaveStatus) error {
	var leave domain.Leave
	err := r.db.First(&leave, id).Error
	if err != nil {
		return err
	}
	leave.Status = status
	return r.db.Save(&leave).Error
}