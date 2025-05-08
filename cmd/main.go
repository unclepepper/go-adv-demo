package main

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)

	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Hello World Server is listening on :8080")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
