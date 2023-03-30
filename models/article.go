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
	Tags          []*Tag `gorm:"many2many:article_tags"`
	FavoriteCount uint
	UserID        uint
}

type Tag struct {
	gorm.Model
	Name     string
	Articles []*Article `gorm:"many2many:article_tags"`
}

func (a *Article) SaveArticle() (*Article, error) {
	err := DB.Create(&a).Error

	if err != nil {
		return &Article{}, err
	}
	return a, nil
}

func ListArticles() ([]Article, error) {
	var articles []Article
	//var tag []Tag
	err := DB.Model(&Article{}).Preload("Tags").Find(&articles).Error
	return articles, err
}
