package http

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-realworld/internal/adapter/storage/repository"
	"github.com/go-realworld/internal/core/domain"
)

var secret string

func init() {
	secret = os.Getenv("API_SECRET")
}

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

	// Check user exists.
	user, err := handler.repository.FindByUsername(loginReq.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			Issuer:    "real-worl-demo-backend",
		}).SignedString(secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Authorization: Token " + token))
}
