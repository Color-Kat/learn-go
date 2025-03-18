package auth

import (
	"demo/http/internal/user"
	"errors"
	"golang.org/x/crypto/bcrypt"
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
		return "", errors.New(ErrUserNotExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &user.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
		Model:    gorm.Model{},
	}

	_, err = service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}

func (service *AuthService) Login(email, password string) (string, error) {
	existingUser, _ := service.UserRepository.FindByEmail(email)
	if existingUser == nil {
		return "", errors.New(ErrWrongCredential)
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredential)
	}

	return existingUser.Email, nil
}
