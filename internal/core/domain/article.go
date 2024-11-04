package domain

import (
	"gorm.io/gorm"
)

// Article is an entity that represents an article.
type Article struct {
	gorm.Model
	Slug          string
	Title         string
	Description   string
	Body          string
	Tags          []*Tag `gorm:"many2many:article_tags"`
	FavoriteCount uint
	UserID        uint
}

// Tag is an entity that represents a tag added to an Article.
type Tag struct {
	gorm.Model
	Name     string
	Articles []*Article `gorm:"many2many:article_tags"`
}
