package auth

import (
	"errors"
	"go/adv-demo/internal/user"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UserRepository.GetByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	u := &user.User{
		Email:    email,
		Password: "",
		Name:     name,
	}
	createdUser, err := service.UserRepository.Create(u)
	if err != nil {
		return "", err
	}
	return createdUser.Email, nil
}
