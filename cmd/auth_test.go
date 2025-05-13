package main

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "test@test.ru",
		Password: "$2a$10$PxkWgzehhJ3izGKoGQNq4.CXnMjdp.vcQhMt/FQW1HSG9GI/pOzeq",
		Name:     "Test User",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "test@test.ru").
		Delete(&user.User{})
}

func TestLoginSuccess(t *testing.T) {
	// Prepare
	db := initDb()
	initData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@test.ru",
		Password: "password",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("got %d, want %d", res.StatusCode, http.StatusOK)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse
	if err := json.Unmarshal(body, &resData); err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatal("got empty token")
	}

	removeData(db)
}

func TestLoginFail(t *testing.T) {
	// Prepare
	db := initDb()
	initData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@test.ru",
		Password: "3",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("got %d, want %d", res.StatusCode, http.StatusUnauthorized)
	}

	removeData(db)
}
