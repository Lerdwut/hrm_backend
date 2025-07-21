package service

import (
	"hr_management/internal/core/domain"
	"hr_management/internal/core/port"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (us *UserService) Register(user *domain.User) (*domain.User, error) {
	user, err := us.repo.CreateUser(user)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, domain.ErrConflictingData
		}
		return nil, domain.ErrInternal
	}
	return user, nil
}

func (us *UserService) ListUsers() ([]*domain.User, error) {
	user, err := us.repo.GetAllUsers()
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, domain.ErrDataNotFound
		}
		return nil, domain.ErrInternal
	}
	return user, nil
}

func (us *UserService) FindByUsername(username string) (*domain.User, error) {
	user, err := us.repo.FindByUsername(username)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, domain.ErrDataNotFound
		}
		return nil, domain.ErrInternal
	}
	return user, nil
}

func (us *UserService) FindByEmail(email string) (*domain.User, error) {
	return us.repo.FindByEmail(email)
}

func (us *UserService) FindByGoogleID(googleID string) (*domain.User, error) {
	return us.repo.FindByGoogleID(googleID)
}

func (us *UserService) CreateOAuthUser(oauthUser *domain.OAuthUser) (*domain.User, error) {
	user := &domain.User{
		Username:   oauthUser.Name,
		Email:      oauthUser.Email,
		Avatar:     oauthUser.Picture,
		GoogleID:   oauthUser.ID,
		Provider:   string(oauthUser.Provider),
		IsVerified: oauthUser.Verified,
	}

	return us.repo.CreateUser(user)
}

func (us *UserService) UpdateOAuthUser(user *domain.User, oauthUser *domain.OAuthUser) (*domain.User, error) {
	user.Username = oauthUser.Name
	user.Email = oauthUser.Email
	user.Avatar = oauthUser.Picture
	user.GoogleID = oauthUser.ID
	user.Provider = string(oauthUser.Provider)
	user.IsVerified = oauthUser.Verified

	return us.repo.UpdateUser(user)
}
