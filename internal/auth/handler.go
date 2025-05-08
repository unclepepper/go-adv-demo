package auth

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/req"
	"go/adv-demo/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}

	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := LoginResponse{
			Token: h.Config.Auth.Secret,
		}
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		res.Json(w, data, http.StatusOK)

		fmt.Println(body)
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := RegisterResponse{
			Token: h.Config.Auth.Secret,
		}
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		res.Json(w, data, http.StatusOK)

		fmt.Println(body)
	}
}
