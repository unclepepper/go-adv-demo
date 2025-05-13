package auth_test

import (
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "test@example.ru",
	}, nil
}
func (repo *MockUserRepository) GetByEmail(email string) (*user.User, error) {
	return nil, nil
}
func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "test@example.ru"
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(initialEmail, "1", "Test user")
	if err != nil {
		t.Fatal(err)
	}

	if email != initialEmail {
		t.Errorf("got email %s, want %s", email, email)
	}
}
