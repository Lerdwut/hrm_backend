package repository

import (
	"hr_management/internal/core/domain"
	"strings"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	if err := r.db.Create(user).
		Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetAllUsers() ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(id string) error {
	if err := r.db.Delete(&domain.User{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByGoogleID(googleID string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("google_id = ?", googleID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}
	return &user, nil
}
