package domain

import (
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
