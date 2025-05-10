package main

import (
	"github.com/joho/godotenv"
	"go/adv-demo/internal/link"
	"go/adv-demo/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&link.Link{}, &user.User{})
	if err != nil {
		return
	}
}
