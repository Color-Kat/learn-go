package auth

import (
	"demo/http/configs"
	"demo/http/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type AuthHandlerDeps struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			response.Json(w, err.Error(), 402)
		}

		// Validation
		if payload.Email == "" || payload.Password == "" {
			response.Json(w, "Email and Password required", 402)
			return
		}

		match, _ := regexp.MatchString(`[\w\.%+\-]+@[\w\.+\-]+\.[A-Za-z]{2,}`, payload.Email)
		if !match {
			response.Json(w, "Wrong Email", 402)
			return
		}

		res := LoginResponse{
			Token: handler.Config.Auth.SecretToken,
		}
		response.Json(w, res, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}
}
