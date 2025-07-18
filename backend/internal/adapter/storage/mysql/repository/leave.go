package repository

import (
	"gorm.io/gorm"
	"hr_management/internal/core/domain"
)

type GormLeaveRepo struct {
	db *gorm.DB
}

func NewGormLeaveRepo(db *gorm.DB) *GormLeaveRepo {
	return &GormLeaveRepo{db: db}
}

func (r *GormLeaveRepo) Create(l *domain.Leave) error {
	return r.db.Create(l).Error
}

func (r *GormLeaveRepo) GetAll() ([]domain.Leave, error) {
	var leaves []domain.Leave
	err := r.db.Find(&leaves).Error
	return leaves, err
}

func (r *GormLeaveRepo) GetByID(id int) (*domain.Leave, error) {
	var leave domain.Leave
	err := r.db.First(&leave, id).Error
	if err != nil {
		return nil, err
	}
	return &leave, nil
}

func (r *GormLeaveRepo) Update(id uint, status domain.LeaveStatus) error {
	var leave domain.Leave
	err := r.db.First(&leave, id).Error
	if err != nil {
		return err
	}
	leave.Status = status
	return r.db.Save(&leave).Error
}

