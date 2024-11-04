package domain

import (
	"gorm.io/gorm"
)

// Tag is an entity that represents a tag added to an Article.
type Tag struct {
	gorm.Model
	Name     string
	Articles []*Article `gorm:"many2many:article_tags"`
}
