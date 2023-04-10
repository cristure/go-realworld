package models

import (
	"errors"
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

func (a *Article) UpdateArticle(newArticle *Article) (*Article, error) {
	if err := DB.First(&a, a.ID).Error; err != nil {
		return nil, errors.New("Article was not found")
	}

	DB.Save(newArticle)
	return newArticle, nil
}

func (a *Article) DeleteArticle() error {
	if err := DB.Delete(a); err != nil {
		return err.Error
	}

	return nil
}

func ListArticles() ([]*Article, error) {
	var articles []*Article
	//var tag []Tag
	err := DB.Model(&Article{}).Preload("Tags").Find(&articles).Error
	return articles, err
}

func FindArticleBySlug(slug string) (*Article, error) {
	var article Article

	if err := DB.First(&article, "slug = ?", slug).Error; err != nil {
		return nil, errors.New("user not found!")
	}

	return &article, nil
}

func FindTagByName(name string) (*Tag, error) {
	var tag Tag

	if err := DB.First(&tag, "name = ?", name).Error; err != nil {
		return nil, errors.New("tag not found")
	}

	return &tag, nil
}
