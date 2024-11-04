package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/go-realworld/internal/core/domain"
)

// Article will perform all the operations with the DB for domain.Article.
type Article struct {
	db *gorm.DB
}

// NewArticle will create a new repository for domain.Article.
func NewArticle(db *gorm.DB) *Article {
	return &Article{db: db}
}

// Create will persist a new domain.Article in the DB.
func (a *Article) Create(article *domain.Article) error {
	err := a.db.Create(article).Error
	if err != nil {
		return fmt.Errorf("failed to create an article: %w", err)
	}

	return nil
}

// FindByID will return the article with matching ID.
func (a *Article) FindByID(id int64) (*domain.Article, error) {
	var article domain.Article
	err := a.db.Find(&article, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find an article with id '%d': %w", id, err)
	}

	return &article, nil
}
