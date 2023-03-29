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
	TagList       []Tag `gorm:"many2many:article_tags"`
	FavoriteCount uint
	UserID        uint
}

type Tag struct {
	gorm.Model
	Name string
}
