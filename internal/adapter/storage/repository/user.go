package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/go-realworld/internal/core/domain"
)

// User will perform all the operations with the DB for domain.Article.
type User struct {
	db *gorm.DB
}

// NewUser will create a new repository for domain.Article.
func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

// Create will persist a new domain.User in the DB.
func (u *User) Create(user *domain.User) error {
	err := u.db.Create(user).Error
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
