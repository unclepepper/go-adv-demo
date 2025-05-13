package jwt_test

import (
	"go/adv-demo/pkg/jwt"
	"testing"
)

func TestJWT_Create(t *testing.T) {
	const email = "test@test.ru"
	jwtService := jwt.NewJWT("/2+XnmJGz1j3ehIVI/5P9kl+CghrE3DcS7rnT+qar5w=")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}

	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal(data)
	}
	if data.Email != email {
		t.Fatalf("email %s is not equal %s", data.Email, email)
	}
}
