package repository

import (
	"gorm.io/gorm"

	"github.com/go-realworld/internal/core/domain"
)

// User will perform all the operations with the DB for domain.Article.
type User struct {
	db       *gorm.DB
	Username string
}

// NewUser will create a new repository for domain.Article.
func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

// Create will persist a new domain.User in the DB.
func (u *User) Create(user *domain.User) error {
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *User) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := u.db.Find(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
