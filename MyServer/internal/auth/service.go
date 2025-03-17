package auth

import (
	"demo/http/internal/user"
	"errors"
	"gorm.io/gorm"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	user := &user.User{
		Email:    email,
		Password: "",
		Name:     name,
		Model:    gorm.Model{},
	}

	_, err := service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}
