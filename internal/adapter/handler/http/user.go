package http

import (
	"fmt"
	"net/http"

	"github.com/go-realworld/internal/adapter/storage/repository"
	"github.com/go-realworld/internal/core/domain"
)

// UserHandler will hold all actions responsible for the domain.User actions.
type UserHandler struct {
	repository *repository.User
}

// NewUserHandler will create an instance of UserHandler.
func NewUserHandler(repository *repository.User) *UserHandler {
	return &UserHandler{repository}
}

// Register will create a new domain.User in the database.
func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	registerReq := struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Email    string `json:"email" validate:"required"`
	}{}
	err := decode(r.Body, &registerReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a user
	if err = handler.repository.Create(
		&domain.User{
			Username: registerReq.Username,
			Password: registerReq.Password,
			Email:    registerReq.Email,
		}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Login will look for the requested user and impersonate it by attaching a JWT token.
func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	loginReq := struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}{}
	err := decode(r.Body, &loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a user
	user, err := handler.repository.FindByUsername(loginReq.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(user)
	w.WriteHeader(http.StatusOK)
}
