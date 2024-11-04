package domain

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User is an entity representing a user.
type User struct {
	gorm.Model
	Username         string     `gorm:"size:255;not null;unique" json:"username"`
	Password         string     `gorm:"size:255;not null;" json:"password"`
	Email            string     `gorm:"size;255 not null;" json:"email"`
	Bio              string     `gorm:"size:255, not null;" json:"bio"`
	Image            string     `gorm:"size:255; nullable" json:"image"`
	FavoriteArticles []*Article `gorm:"many2many:favorite_article_user"`
	Articles         []Article
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}
