package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Slug          string
	Title         string
	Description   string
	Body          string
	TagList       []string
	FavoriteCount uint
	AuthorID      uint
	User          User `gorm:"foreignKey:AuthorID"`
}
